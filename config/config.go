// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"bytes"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func GetTestRunConfig(fileName string) (TestRunConfig, error) {
	runConfig := TestRunConfig{}

	yamlFile, err := os.Open(fileName)
	if err != nil {
		return runConfig, err
	}
	defer yamlFile.Close()

	byteValue, _ := ioutil.ReadAll(yamlFile)
	tpl, err := template.New("").Option("missingkey=zero").Parse(string(byteValue))
	if err != nil {
		return runConfig, err
	}

	environmentMap := make(map[string]string)
	for _, v := range os.Environ() {
		separated := strings.Split(v, "=")
		environmentMap[separated[0]] = separated[1]
	}

	var writer bytes.Buffer
	err = tpl.Execute(&writer, environmentMap)
	if err != nil {
		return runConfig, err
	}

	err = yaml.Unmarshal(writer.Bytes(), &runConfig)
	if err != nil {
		return runConfig, err
	}

	return runConfig, nil
}

type TestRunConfig struct {
	RunId       string `yaml:"runid"`
	Environment string `yaml:"environment"`
	Project     string `yaml:"project"`
	Link        string `yaml:"link"`
	Owner       string `yaml:"owner"`
	Runner      string `yaml:"runner"`
	StartedBy   string `yaml:"startedby"`
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
