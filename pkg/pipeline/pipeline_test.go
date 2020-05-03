package pipeline

import (
	"testing"

	"github.com/taskctl/taskctl/pkg/config"
)

func TestBuildPipeline_Cyclic(t *testing.T) {
	stages := []*config.StageDefinition{
		{
			Name:      "task1",
			Task:      "task1",
			DependsOn: []string{"last-stage"},
		},
		{
			Name:      "task2",
			Task:      "task2",
			DependsOn: []string{"task1"},
			Env:       nil,
		},
		{
			Name:      "last-stage",
			Task:      "task3",
			DependsOn: []string{"task2"},
		},
	}

	tasks := map[string]*config.TaskDefinition{
		"task1": {
			Name: "task1",
		},
		"task2": {
			Name: "task2",
		},
		"task3": {
			Name: "task3",
		},
	}

	_, err := BuildPipeline(stages, make(map[string][]*config.StageDefinition), tasks)
	if err == nil || err.Error() != "cycle detected" {
		t.Errorf("cycles detection failed")
	}
}
