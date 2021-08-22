package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigInfo struct {
	Port       string  `yaml:"port"`
	ApiVersion string  `yaml:"apiVersion"`
	LogLevel   int     `yaml:"logLevel"`
	MongoDb    MongoDb `yaml:"mongoDb"`
}

type MongoDb struct {
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	MaxPoolSize       uint64 `yaml:"maxPoolSize"`
	ConnectCtxTimeout int    `yaml:"connectCtxTimeout"`
	OperateCtxTimeout int    `yaml:"operateCtxTimeout"`
}

var Config = new(ConfigInfo)

func init() {
	//testing.Init()
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
