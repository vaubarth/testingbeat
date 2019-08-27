// +build !integration

package config

import (
	td "github.com/maxatome/go-testdeep"
	"testing"
)

var expectedRunConfig = TestRunConfig{
	RunId:       "runid",
	Environment: "env",
	Project:     "project",
	Link:        "link",
	Owner:       "owner",
	Runner:      "runner",
	StartedBy:   "startedBy",
}

func TestGetTestRunConfig(t *testing.T) {
	runConfig, _ := GetTestRunConfig("../tests/runconfig.yml")
	td.Cmp(t, runConfig, expectedRunConfig)
}
