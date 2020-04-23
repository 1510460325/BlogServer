package config

import (
	"blog/util/fileUtil"
	"gopkg.in/yaml.v2"
)

var AppSetting = &appSetting{}

type appSetting struct {
	DB struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		Timeout  int    `yaml:"timeout"`
	} `yaml:"db"`
	Log struct {
		Path     string `yaml:"path"`
		Filename string `yaml:"filename"`
		Level    string `yaml:"level"`
	} `yaml:"log"`
	Server struct{
		Port uint32 `yaml:"port"`
	} `yaml:"server"`
}

/**
 * 自动加载配置文件
 */
func init() {
	data := fileUtil.ReadFile("config.yml")
	err := yaml.Unmarshal([]byte(data), AppSetting)
	if err != nil {
		panic(err)
	}
}
