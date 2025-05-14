package validate

import (
	"testing"
)

func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
		errMsg    string
	}{
		{
			name:      "valid_simple_name",
			input:     "Project1",
			expectErr: false,
		},
		{
			name:      "empty_name",
			input:     "",
			expectErr: true,
			errMsg:    "empty project name",
		},
		{
			name:      "too_long_name",
			input:     "ThisIsAVeryLongProjectNameThatExceedsTheAllowedSixtyFourCharactersLimit",
			expectErr: true,
			errMsg:    "too long",
		},
		{
			name:      "invalid_start_digit",
			input:     "1InvalidName",
			expectErr: true,
			errMsg:    "invalid characters",
		},
		{
			name:      "invalid_special_character",
			input:     "Invalid#Name!",
			expectErr: true,
			errMsg:    "invalid characters",
		},
		{
			name:      "valid_name_with_dash_and_underscore",
			input:     "valid-name_1",
			expectErr: false,
		},
		{
			name:      "reserved_keyword",
			input:     "package",
			expectErr: true,
			errMsg:    "'package' — reserver Go keyword",
		},
		{
			name:      "another_reserved_keyword",
			input:     "func",
			expectErr: true,
			errMsg:    "'func' — reserver Go keyword",
		},
		{
			name:      "valid_name_all_caps",
			input:     "PROJECT_AAA",
			expectErr: false,
		},
		{
			name:      "valid_name_with_only_one_char",
			input:     "X",
			expectErr: false,
		},
		{
			name:      "valid_name_starting_with_underscore",
			input:     "_InvalidName123",
			expectErr: true,
			errMsg:    "invalid characters",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateProjectName(tc.input)
			if tc.expectErr {
				if err == nil {
					t.Errorf("expected an error but got nil")
				} else if err.Error() != tc.errMsg {
					t.Errorf("expected error message '%s', but got '%s'", tc.errMsg, err.Error())
				}
			} else if err != nil {
				t.Errorf("expected no error, but got: %s", err.Error())
			}
		})
	}
}
