// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bearertokenauthextension

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		id          config.ComponentID
		expected    config.Extension
		expectedErr bool
	}{
		{
			id:          config.NewComponentID(typeStr),
			expectedErr: true,
		},
		{
			id: config.NewComponentIDWithName(typeStr, "sometoken"),
			expected: &Config{
				ExtensionSettings: config.NewExtensionSettings(config.NewComponentID(typeStr)),
				Scheme:            defaultScheme,
				BearerToken:       "sometoken",
			},
		},
		{
			id: config.NewComponentIDWithName(typeStr, "withscheme"),
			expected: &Config{
				ExtensionSettings: config.NewExtensionSettings(config.NewComponentID(typeStr)),
				Scheme:            "MyScheme",
				BearerToken:       "my-token",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
			require.NoError(t, err)
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()
			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, config.UnmarshalExtension(sub, cfg))
			if tt.expectedErr {
				assert.Error(t, cfg.Validate())
				return
			}
			assert.NoError(t, cfg.Validate())
			assert.Equal(t, tt.expected, cfg)
		})
	}
}