// Copyright ServiceNow, Inc
// SPDX-License-Identifier: Apache-2.0

package otelsampling

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseCloudObsTraceState(t *testing.T) {
	type testCase struct {
		in        string
		sval      string
		expectErr error
	}
	const ns = ""
	for _, test := range []testCase{
		// s-value correct cases
		{"s:2", "2", nil},
		{"s:123", "123", nil},

		// syntax errors
		{"", ns, strconv.ErrSyntax},
		{"t=1,", ns, strconv.ErrSyntax},
		{"s:-1", ns, strconv.ErrSyntax},
	} {
		t.Run(testName(test.in), func(t *testing.T) {
			otts, err := NewCloudObsTraceState(test.in)

			if test.expectErr != nil {
				require.True(t, errors.Is(err, test.expectErr), "%q: not expecting %v wanted %v", test.in, err, test.expectErr)
				return
			}

			require.NoError(t, err)
			if test.sval != ns {
				require.NotEqual(t, "", otts.SValue())
				require.Equal(t, test.sval, otts.SValue())
			} else {
				require.Equal(t, "", otts.SValue(), "should have no s-value: %s", otts.SValue())
			}
			// on success Serialize() should not modify
			// test by re-parsing
			var w strings.Builder
			otts.Serialize(&w)
			cpy, err := NewCloudObsTraceState(w.String())
			require.NoError(t, err)
			require.Equal(t, otts, cpy)
		})
	}
}

func TestUpdateSValueWithSampling(t *testing.T) {
	type testCase struct {
		// The input otel tracestate; no error conditions tested
		in string

		// The incoming adjusted count; defined whether
		// s-value is present or not.
		adjCountIn float64

		// the update probability; threshold and tvalue are
		// derived from this
		prob float64

		// when update error is expected
		updateErr error

		// output s-value
		out string

		// output adjusted count
		adjCountOut float64
	}
	for _, test := range []testCase{
		// 8/16 in, 2/16 out
		{"s:8", 2, 0x2p-4, nil, "s:2", 8},

		// 1/16 in, 50% update (error)
		{"s:1", 16, 0x8p-4, ErrInconsistentSampling, "s:1", 16},

		// no sampling in, 1/16 update
		{"", 0, 0x1p-4, nil, "s:1", 16},

		// none in, 100% update
		{"", 0, 1, nil, "", 0},

		// 1/2 in, 100% update (error)
		{"s:8", 2, 1, ErrInconsistentSampling, "s:8", 2},

		// a/16 in, 5/16 out
		{"s:a", 16.0 / 10, 0x5 / 16.0, nil, "s:5", 16.0 / 5},
	} {
		t.Run(test.in+"/"+test.out, func(t *testing.T) {
			cots := CloudObsTraceState{}
			if test.in != "" {
				var err error
				cots, err = NewCloudObsTraceState(test.in)
				require.NoError(t, err)
			}

			require.Equal(t, test.adjCountIn, cots.AdjustedCount())

			newTh, err := ProbabilityToThreshold(test.prob)
			require.NoError(t, err)

			upErr := cots.UpdateSValueWithSampling(newTh)

			if test.updateErr != nil {
				require.Equal(t, test.updateErr, upErr)
			}

			var outData strings.Builder
			err = cots.Serialize(&outData)
			require.NoError(t, err)
			require.Equal(t, test.out, outData.String())

			require.Equal(t, test.adjCountOut, cots.AdjustedCount())
		})
	}
}
