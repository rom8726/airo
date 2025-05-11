package generator

import (
	"context"
	"fmt"
	"os"

	"github.com/rom8726/airo/config"
)

func GenerateProject(ctx context.Context, cfg *config.ProjectConfig) error {
	fmt.Println("Generate project directory...")
	root := cfg.ProjectName
	if err := os.MkdirAll(root, 0755); err != nil {
		return err
	}
	fmt.Println("project directory: done")

	fmt.Println("Write go.mod...")
	// 1. go.mod
	if err := writeGoMod(cfg); err != nil {
		return err
	}
	fmt.Println("go.mod: done")

	fmt.Println("Generate OpenAPI server with ogen...")
	// 2. ogen
	if err := runOGen(cfg); err != nil {
		return fmt.Errorf("ogen failed: %w", err)
	}
	fmt.Println("ogen: done")

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
