// Copyright ServiceNow, Inc
// SPDX-License-Identifier: Apache-2.0

package satellitesamplerprocessor

import (
	"errors"
	"fmt"
)

var (
	errInvalidPercent = errors.New("sampling percent must be in (0, 100]")
)

// Config defines configuration for Lightstep's "classic" Satellite sampler.
type Config struct {
	// Percent in the range (0, 100].  Defaults to 100.
	//
	// Note that satellites began supporting percent-based
	// sampling configuration at release 2022-04-28_17-39-22Z.
	// When OneInN is set instead, use the formula `100.0 /
	// float64(OneInN)`.
	Percent float64 `mapstructure:"percent"`
}

func (c *Config) Validate() error {
	if c.Percent <= 0 || c.Percent > 100 {
		return fmt.Errorf("%w: %v", errInvalidPercent, c.Percent)
	}
	return nil
}
