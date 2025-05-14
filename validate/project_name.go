package validate

import (
	"errors"
	"fmt"
	"go/token"
	"regexp"
)

func ValidateProjectName(name string) error {
	if name == "" {
		return errors.New("empty project name")
	}
	if len(name) > 64 {
		return errors.New("too long")
	}

	var re = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_-]*$`)
	if !re.MatchString(name) {
		return errors.New("invalid characters")
	}

	if token.Lookup(name).IsKeyword() {
		return fmt.Errorf("'%s' — reserver Go keyword", name)
	}

	return nil
}
