package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigInfo struct {
	Port       string    `yaml:"port"`
	ApiVersion string    `yaml:"apiVersion"`
	LogLevel   int `yaml:"logLevel"`
}

var Config = new(ConfigInfo)

func init() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "config_files/config.yaml", "config file path")
	flag.Parse()

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
