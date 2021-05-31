package config

import (
	"embed"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
}

var config *Config
var RunMode string

//go:embed *
var configFile embed.FS

// 默认配置
func init() {
	runmode := os.Getenv("RUN_MODE")
	if runmode == "" {
		runmode = RunMode
	}
	configPath := "config.dev.yml"
	if runmode == "product" {
		configPath = "config.yml"
	}
	log.Println("当前配置:", configPath)
	body, err := configFile.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(body, &config)
	if err != nil {
		log.Fatalln(err)
	}
}

// Get 获取配置文件
func Get() *Config {
	return config
}
