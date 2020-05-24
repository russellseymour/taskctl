package config

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/taskctl/taskctl/pkg/task"

	"gopkg.in/yaml.v2"
)

var testConfig, _ = ioutil.ReadFile("testdata/tasks.yaml")

func TestConfig_decode(t *testing.T) {
	loader := NewConfigLoader()

	var cm = make(map[string]interface{})
	var dec = yaml.NewDecoder(bytes.NewReader(testConfig))
	dec.SetStrict(true)

	err := dec.Decode(cm)
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := loader.decode(cm)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := cfg.Tasks["test-task"]; !ok {
		t.Fatal("tasks parsing error")
	}

	if _, ok := cfg.Pipelines["pipeline2"]; !ok {
		t.Fatal("pipelines parsing error")
	}

	if len(cfg.Pipelines["pipeline2"]) != 2 {
		t.Fatal("pipelines parsing failed")
	}
}

func TestConfig_merge(t *testing.T) {
	cfg1 := &Config{
		Tasks: map[string]*task.Task{"task1": {}},
	}

	cfg2 := &Config{
		Tasks: map[string]*task.Task{"task2": {}},
	}

	err := cfg1.merge(cfg2)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := cfg1.Tasks["task2"]; !ok {
		t.Error()
	}
}
