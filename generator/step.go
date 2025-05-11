package generator

import (
	"context"

	"github.com/rom8726/airo/config"
)

type Step interface {
	Description() string
	Do(ctx context.Context, cfg *config.ProjectConfig) error
}
