// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

type TestRunConfig struct {
	RunId       string `yaml:"run_id"`
	Environment string `yaml:"environment"`
	Project     string `yaml:"project"`
	Link        string `yaml:"link"`
	Owner       string `yaml:"owner"`
	Runner      string `yaml:"runner"`
	StartedBy   string `yaml:"started_by"`
}

func GetTestRunConfig(fileName string) (TestRunConfig, error) {
	runConfig := TestRunConfig{}

	yamlFile, err := os.Open(fileName)
	if err != nil {
		return runConfig, err
	}
	defer yamlFile.Close()
	byteValue, _ := ioutil.ReadAll(yamlFile)

	err = yaml.Unmarshal(byteValue, &runConfig)
	if err != nil {
		return runConfig, err
	}

	return runConfig, nil
}

type Config struct {
	Path          string `config:"path"`
	Type          string `config:"type"`
	RunConfigFile string `config:"runconfig"`
}

var DefaultConfig = Config{
	Path: ".",
	Type: "junit",
}
