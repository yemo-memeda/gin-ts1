package global

import (
	"github.com/spf13/viper"
	"github.com/yemo-memeda/gin-ts1/config"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
}

var App = new(Application)
