package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dollarkillerx/env"
	"gopkg.in/yaml.v2"
)

type BaseConfig struct {
	SocketAddr  string `yaml:"SocketAddr"`
	Username    string `yaml:"Username"`
	Password    string `yaml:"Password"`
	StoragePath string `yaml:"StoragePath"`

	Debug bool `yaml:"Debug"`
}

// 初始化Config
func InitConfig() *BaseConfig {
	var conf BaseConfig
	cfg, err := ioutil.ReadFile("./configs/config.yaml")
	if err != nil {
		if err := env.Fill(&conf); err == nil {
			return &conf
		}
		log.Fatalln(fmt.Errorf("InitConfig Error: %s", err.Error()))
	}
	if err := yaml.Unmarshal(cfg, &conf); err != nil {
		log.Fatalln(fmt.Errorf("InitConfig Error: %s", err.Error()))
	}
	return &conf
}
