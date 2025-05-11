package generator

import (
	"context"
	"fmt"
	"os"

	"github.com/rom8726/airo/config"
)

func GenerateProject(ctx context.Context, cfg *config.ProjectConfig) error {
	root := cfg.ProjectName

	if err := os.MkdirAll(root, 0755); err != nil {
		return err
	}

	//// 1. go.mod
	//if err := writeGoMod(cfg); err != nil {
	//	return err
	//}

	// 2. ogen
	if err := runOGen(cfg); err != nil {
		return fmt.Errorf("ogen failed: %w", err)
	}

	//// 3. main.go, app.go, config.go
	//if err := writeMain(cfg); err != nil {
	//	return err
	//}
	//if err := writeConfig(cfg); err != nil {
	//	return err
	//}
	//if err := writeApp(cfg); err != nil {
	//	return err
	//}
	//
	//// 4. infra (postgres / redis)
	//if cfg.UsePostgres {
	//	if err := writePostgres(cfg); err != nil {
	//		return err
	//	}
	//}
	//if cfg.UseRedis {
	//	if err := writeRedis(cfg); err != nil {
	//		return err
	//	}
	//}

	return nil
}
