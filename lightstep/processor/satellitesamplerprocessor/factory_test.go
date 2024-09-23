// Copyright ServiceNow, Inc
// SPDX-License-Identifier: Apache-2.0

package satellitesamplerprocessor

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
)

func TestDefaultConfig(t *testing.T) {
	require.Equal(t, component.Config(testConfig(100)), createDefaultConfig())
}
