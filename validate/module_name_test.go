package validate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateModuleName(t *testing.T) {
	tests := []struct {
		name       string
		moduleName string
		wantErr    bool
	}{
		{
			name:       "Valid module name",
			moduleName: "github.com/example/module",
			wantErr:    false,
		},
		{
			name:       "Empty name",
			moduleName: "",
			wantErr:    true,
		},
		{
			name:       "Too long name",
			moduleName: "a12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			wantErr:    true,
		},
		{
			name:       "Invalid characters",
			moduleName: "github.com/exa mple/module",
			wantErr:    true,
		},
		{
			name:       "Relative path",
			moduleName: "./relative/path",
			wantErr:    true,
		},
		{
			name:       "Non-ASCII characters",
			moduleName: "github.com/example/模块",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateModuleName(tt.moduleName)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
