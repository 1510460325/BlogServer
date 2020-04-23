package logUtil

import (
	"blog/config"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

var Logger *logrus.Logger = nil

func init() {
	fileName := path.Join(config.AppSetting.Log.Path, config.AppSetting.Log.Filename)
	if _, err := os.Open(config.AppSetting.Log.Path); err != nil {
		if err := os.Mkdir(config.AppSetting.Log.Path, os.ModeDir); err != nil {
			panic(err)
		}
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		src, err = os.Create(fileName)
	}
	Logger = logrus.New()
	Logger.SetOutput(src)
	level, _ := logrus.ParseLevel(config.AppSetting.Log.Level)
	Logger.SetLevel(level)
	Logger.SetFormatter(&logrus.TextFormatter{})
}
