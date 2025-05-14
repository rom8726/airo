package validate

import (
	"errors"
	"testing"

	"golang.org/x/mod/module"
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
			name:       "Reserved keywords",
			moduleName: "github.com/example/.git",
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
			var errInvalidPath *module.InvalidPathError

			err := ValidateModuleName(tt.moduleName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateModuleName(%q) = %v, wantErr %v", tt.moduleName, err, tt.wantErr)
			} else if err != nil && !tt.wantErr {
				t.Errorf("Unexpected error: %v", err)
			} else if err != nil && tt.wantErr && !errors.As(err, &errInvalidPath) && tt.moduleName != "" {
				t.Errorf("Error type mismatch: got %v, expected module.ErrInvalidPath or related", err)
			}
		})
	}
}
