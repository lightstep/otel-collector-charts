// Copyright ServiceNow, Inc
// SPDX-License-Identifier: Apache-2.0

package otelsampling

import (
	"io"
	"strconv"
)

// This file borrows everything it can from oteltracestate.go.  The
// s-value here uses a different representation from the OTel T-value,
// because it was adopted before the encoding finalized in OTEP 235.
//
// The Satellite's S-value is an acceptance threshold, vs the OTel
// T-value which is a rejection thresold.  There is a 1:1 mapping.

type CloudObsTraceState struct {
	commonTraceState

	// s-value and threshold are represented the same as t-value
	threshold Threshold // t value parsed, as a threshold
	svalue    string    // 1-14 ASCII hex digits
}

const (
	// RName is the CloudObs tracestate field for S-value, which
	// uses the same encoding as t-value.
	SName = "s"
)

func NewCloudObsTraceState(input string) (cots CloudObsTraceState, _ error) {
	if len(input) > hardMaxOTelLength {
		return cots, ErrTraceStateSize
	}

	if !otelTracestateRe.MatchString(input) {
		return CloudObsTraceState{}, strconv.ErrSyntax
	}

	err := otelSyntax.scanKeyValues(input, func(key, value string) error {
		switch key {
		case SName:
			// S-value is represented as (MaxAdjustedCount - TValue)
			// because it is an acceptance threshold vs a rejection
			// threshold.
			tv, err := TValueToThreshold(value)
			if err != nil {
				cots.threshold = AlwaysSampleThreshold
				return err
			}
			// We do not expect tv.unsigned == 0, which is zero
			// adjusted count, from a satellite sampler.
			if tv.unsigned == 0 {
				cots.threshold = NeverSampleThreshold
			} else {
				cots.svalue = value
				cots.threshold, _ = UnsignedToThreshold(MaxAdjustedCount - tv.Unsigned())
			}
		default:
			cots.kvs = append(cots.kvs, KV{
				Key:   key,
				Value: value,
			})
		}
		return nil
	})

	return cots, err
}

func (cots *CloudObsTraceState) SValue() string {
	return cots.svalue
}

func (cots *CloudObsTraceState) SValueThreshold() (Threshold, bool) {
	return cots.threshold, cots.svalue != ""
}

func (cots *CloudObsTraceState) UpdateSValueWithSampling(sampledThreshold Threshold) error {
	if len(cots.SValue()) != 0 && ThresholdGreater(cots.threshold, sampledThreshold) {
		return ErrInconsistentSampling
	}
	cots.threshold = sampledThreshold

	inv, _ := UnsignedToThreshold(MaxAdjustedCount - sampledThreshold.Unsigned())
	cots.svalue = inv.TValue()
	return nil
}

func (cots *CloudObsTraceState) AdjustedCount() float64 {
	if len(cots.svalue) == 0 {
		return 0
	}
	return cots.threshold.AdjustedCount()
}

func (cots *CloudObsTraceState) ClearSValue() {
	cots.svalue = ""
	cots.threshold = AlwaysSampleThreshold
}

func (cots *CloudObsTraceState) HasAnyValue() bool {
	return len(cots.SValue()) != 0 || len(cots.ExtraValues()) != 0
}

func (cots *CloudObsTraceState) Serialize(w io.StringWriter) error {
	ser := serializer{writer: w}
	cnt := 0
	sep := func() {
		if cnt != 0 {
			ser.write(";")
		}
		cnt++
	}
	if len(cots.SValue()) != 0 {
		sep()
		ser.write(SName)
		ser.write(":")
		ser.write(cots.SValue())
	}
	for _, kv := range cots.ExtraValues() {
		sep()
		ser.write(kv.Key)
		ser.write(":")
		ser.write(kv.Value)
	}
	return ser.err
}
