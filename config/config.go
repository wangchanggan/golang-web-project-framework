package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigInfo struct {
	Port       string `yaml:"port"`
	ApiVersion string `yaml:"apiVersion"`
}

var Config = new(ConfigInfo)

func init() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config/config.yaml", "config file path")
	flag.Parse()
	fmt.Println(configFilePath)

	if configFilePath != "" {
		configBytes, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(configBytes, &Config)
		if err != nil {
			panic(err)
		}
	} else {
		flag.Usage()
		panic("lack of config file")
	}
}
