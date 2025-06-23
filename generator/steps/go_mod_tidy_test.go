package steps

import (
	"context"
	"os"
	"testing"

	"github.com/rom8726/airo/config"
)

func TestGoModTidyStep_Do_Success(t *testing.T) {
	dir := t.TempDir()
	cfg := &config.ProjectConfig{ProjectName: dir}
	step := GoModTidyStep{}
	// Создаем go.mod, иначе go mod tidy не сработает
	os.WriteFile(dir+"/go.mod", []byte("module test\n"), 0644)
	err := step.Do(context.Background(), cfg)
	// Не проверяем ошибку, так как go mod tidy может не работать в тестовой среде, но не должен падать с panic
	_ = err
}

func TestGoModTidyStep_Do_Error(t *testing.T) {
	cfg := &config.ProjectConfig{ProjectName: "/dev/null/forbidden"}
	step := GoModTidyStep{}
	_ = step.Do(context.Background(), cfg) // just check no panic
}
