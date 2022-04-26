package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	UserServer struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	}
	GoodsServer struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	}
	Mysql struct {
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Prefix   string `yaml:"prefix"`
	}
}

func GetConfig(filePath string) Config {
	config := Config{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}

	fmt.Println(string(content))
	fmt.Printf("init data: %v", config)
	if yaml.Unmarshal(content, &config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	fmt.Printf("File config: %v", config)
	return config
}
