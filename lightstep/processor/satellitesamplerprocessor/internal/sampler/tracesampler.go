package sampler

import (
	"hash/crc32"
	"math"
)

var SampleEverything = TraceSampler{
	partition: math.MaxUint32,
	rate:      1,
}

// TraceSampler is a uint64 where all spans with trace id <= TraceSampler will be included in the sample.
// All spans with trace id > partition will NOT be included in the sample.
// This is a separate type to avoid missing conversion of an old one-in-n sample rate to use the new sample method.
type TraceSampler struct {
	partition uint32
	// Rate is the one-in-n-sample rate.
	// This is used in many parts of the microsat for adjusting counts.
	// Rather than recomputing it repeatedly, it's more economical to compute it once for the sampler.
	rate float64
}

// NewTraceSampler simply returns a new sampler, for use outside the
// Satellite code base.  Satellites use ComputeTraceSampler.
func NewTraceSampler(samplePercent float64) TraceSampler {
	part, rate := func() (uint32, float64) {
		switch samplePercent {
		case 100:
			return math.MaxUint32, 1
		case 0:
			return 0, 0
		default:
			floatPartition := float64(math.MaxUint32) * samplePercent / 100
			floatRate := 1 / (floatPartition / float64(math.MaxUint32))
			return uint32(floatPartition), floatRate
		}
	}()

	return TraceSampler{partition: part, rate: rate}
}

// IsSampledOut returns true when a span should NOT pass through the
// pipeline.
func (s TraceSampler) IsSampledOut(traceID string) bool {
	if s.partition == math.MaxUint32 {
		return false
	}
	if s.partition == 0 {
		return true
	}

	hash := crc32.Checksum([]byte(traceID), crc32.IEEETable)
	return hash > s.partition
}

func (s TraceSampler) GetSampleRate() float64 {
	return s.rate
}
