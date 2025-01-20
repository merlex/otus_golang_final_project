package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Cache  *CacheConf
	Logger *LoggerConf
	HTTP   *HTTPServerConfig
}

type CacheConf struct {
	Dir      string `default:"/tmp/" yaml:"dir"`
	Capacity int    `default:"10" yaml:"capacity"`
}

type LoggerConf struct {
	Level        string `default:"info" yaml:"level"`
	File         string `default:"image_previewer_proxy.log" yaml:"file"`
	LogToFile    bool   `split_words:"true" default:"false" yaml:"logToFile"`
	LogToConsole bool   `split_words:"true" default:"true" yaml:"logToConsole"`
}

type HTTPServerConfig struct {
	IP   string `default:"0.0.0.0" yaml:"ip"`
	Port string `default:"8585" yaml:"port"`
}

func New() *Config {
	return &Config{
		Cache:  &CacheConf{Dir: "/tmp/", Capacity: 10},
		Logger: &LoggerConf{Level: "info"},
		HTTP:   &HTTPServerConfig{IP: "0.0.0.0", Port: "8585"},
	}
}

func ReadConfig(configFile string) *Config {
	conf := &Config{}
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Could not open a config file %v\n", err)
	} else {
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			fmt.Printf("Could not unmarshal a config file %v\n", err)
		}
	}
	if conf.Logger == nil {
		err = envconfig.Process("", conf)
		if err != nil {
			fmt.Printf("Process ENV variables %v\n", err)
			fmt.Println("Will use default config")
			conf = New()
		}
	}
	return conf
}
