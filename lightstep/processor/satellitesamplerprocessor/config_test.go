// Copyright ServiceNow, Inc
// SPDX-License-Identifier: Apache-2.0

package satellitesamplerprocessor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testConfig(pct float64) *Config {
	return &Config{Percent: pct}
}

func TestConfigValidate(t *testing.T) {
	require.Error(t, testConfig(-1).Validate())
	require.Error(t, testConfig(0).Validate())
	require.Error(t, testConfig(101).Validate())

	require.NoError(t, testConfig(1).Validate())
	require.NoError(t, testConfig(50).Validate())
	require.NoError(t, testConfig(99).Validate())
	require.NoError(t, testConfig(100).Validate())
}
