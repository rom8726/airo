package validate

import (
	"errors"

	"golang.org/x/mod/module"
)

func ValidateModuleName(name string) error {
	if name == "" {
		return errors.New("empty module name")
	}
	if len(name) > 128 {
		return errors.New("too long")
	}

	if err := module.CheckImportPath(name); err != nil {
		return err
	}

	return nil
}
