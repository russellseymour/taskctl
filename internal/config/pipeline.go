package config

import (
	"fmt"

	"github.com/Ensono/taskctl/pkg/variables"

	"github.com/Ensono/taskctl/pkg/scheduler"
	"github.com/Ensono/taskctl/pkg/task"
)

func buildPipeline(g *scheduler.ExecutionGraph, stages []*PipelineDefinition, cfg *Config) (*scheduler.ExecutionGraph, error) {
	for _, def := range stages {
		var stageTask *task.Task
		var stagePipeline *scheduler.ExecutionGraph

		if def.Task != "" {
			stageTask = cfg.Tasks[def.Task]
			if stageTask == nil {
				return nil, fmt.Errorf("stage build failed: no such task %s", def.Task)
			}
		} else {
			stagePipeline = cfg.Pipelines[def.Pipeline]
			if stagePipeline == nil {
				return nil, fmt.Errorf("stage build failed: no such pipeline %s", def.Task)
			}
		}

		stage := scheduler.NewStage(func(s *scheduler.Stage) {
			s.Name = def.Name
			s.Condition = def.Condition
			s.Task = stageTask
			s.Pipeline = stagePipeline
			s.DependsOn = def.DependsOn
			s.Dir = def.Dir
			s.AllowFailure = def.AllowFailure
			s.Env = variables.FromMap(def.Env)
			s.Variables = variables.FromMap(def.Variables)
		})

		if stage.Dir != "" {
			stage.Task.Dir = stage.Dir
		}

		if stage.Name == "" {
			if def.Task != "" {
				stage.Name = def.Task
			}

			if def.Pipeline != "" {
				stage.Name = def.Pipeline
			}

			if stage.Name == "" {
				return nil, fmt.Errorf("stage for task %s must have name", def.Task)
			}
		}

		stage.Variables.Set(".Stage.Name", stage.Name)

		if _, err := g.Node(stage.Name); err == nil {
			return nil, fmt.Errorf("stage with same name %s already exists", stage.Name)
		}

		err := g.AddStage(stage)
		if err != nil {
			return nil, err
		}
	}

	if err := g.HasCycle(); err != nil {
		return nil, err
	}

	return g, nil
}
