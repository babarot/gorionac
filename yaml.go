package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

const yamlFile = "package.yaml"

type Yaml struct {
	Package []Package `yaml:"package"`
	Failed  []string
}

type Package struct {
	Name      string `yaml:"name"`
	Noupdate  bool   `yaml:"noupdate"`
	Noinstall bool   `yaml:"noinstall"`
}

func readYaml() (y Yaml, err error) {
	if _, err := os.Stat(yamlFile); err != nil {
		return y, fmt.Errorf("%s: no such file or directory", yamlFile)
	}

	buf, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf, &y)
	if err != nil {
		str := []byte(err.Error())
		assigned := regexp.MustCompile(`(line \d+)`)
		group := assigned.FindSubmatch(str)
		if len(group) != 0 {
			err = fmt.Errorf("Syntax Error at %s in %s", string(group[0]), yamlFile)
		}
	}

	return
}
