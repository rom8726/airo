package generator

import (
	"context"
	"errors"
	"testing"

	"github.com/rom8726/airo/config"
	"github.com/rom8726/airo/generator/infra"

	"github.com/stretchr/testify/require"
)

type mockStep struct {
	desc     string
	doCalled bool
	fail     bool
}

func (m *mockStep) Description() string { return m.desc }
func (m *mockStep) Do(_ context.Context, _ *config.ProjectConfig) error {
	m.doCalled = true
	if m.fail {
		return errors.New("fail")
	}
	return nil
}

func TestGenerator_NewAndWithSteps(t *testing.T) {
	reg := infra.NewRegistry()
	g := New(reg)
	steps := []StepProvider{
		func(_ *infra.Registry) Step { return &mockStep{desc: "step1"} },
		func(_ *infra.Registry) Step { return &mockStep{desc: "step2"} },
	}
	g.WithSteps(steps)
	require.Len(t, g.stepProviders, 2)
}

func TestGenerator_AddStep(t *testing.T) {
	reg := infra.NewRegistry()
	g := New(reg)
	initLen := len(g.stepProviders)
	g.AddStep(func(_ *infra.Registry) Step { return &mockStep{desc: "added"} })
	require.Equal(t, initLen+1, len(g.stepProviders))
}

func TestGenerator_GenerateProject_AllStepsOK(t *testing.T) {
	reg := infra.NewRegistry()
	g := New(reg).WithSteps([]StepProvider{
		func(_ *infra.Registry) Step { return &mockStep{desc: "ok1"} },
		func(_ *infra.Registry) Step { return &mockStep{desc: "ok2"} },
	})
	cfg := &config.ProjectConfig{}
	err := g.GenerateProject(context.Background(), cfg)
	require.NoError(t, err)
}

func TestGenerator_GenerateProject_StepFails(t *testing.T) {
	reg := infra.NewRegistry()
	failStep := &mockStep{desc: "fail", fail: true}
	g := New(reg).WithSteps([]StepProvider{
		func(_ *infra.Registry) Step { return &mockStep{desc: "ok"} },
		func(_ *infra.Registry) Step { return failStep },
		func(_ *infra.Registry) Step { return &mockStep{desc: "never"} },
	})
	cfg := &config.ProjectConfig{}
	err := g.GenerateProject(context.Background(), cfg)
	require.Error(t, err)
	require.True(t, failStep.doCalled)
}

func TestGenerator_GenerateProject_ContextCancelled(t *testing.T) {
	reg := infra.NewRegistry()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	g := New(reg).WithSteps([]StepProvider{
		func(_ *infra.Registry) Step { return &mockStep{desc: "should not run"} },
	})
	cfg := &config.ProjectConfig{}
	err := g.GenerateProject(ctx, cfg)
	require.Error(t, err)
}
