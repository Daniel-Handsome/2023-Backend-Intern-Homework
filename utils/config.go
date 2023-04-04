package utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var App *Config

type Config struct {
	DB `mapstructure:",squash"`
}

type DB struct {
	Connection             string        `mapstructure:"DB_CONNECTION"`
	Host                   string        `mapstructure:"DB_HOST"`
	Port                   int32         `mapstructure:"DB_PORT"`
	Database               string        `mapstructure:"DB_DATABASE"`
	Username               string        `mapstructure:"DB_USERNAME"`
	Password               string        `mapstructure:"DB_PASSWORD"`
	Token_symmetric_key    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	Access_token_duration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	Refresh_token_duration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	Grpc_port              int32         `mapstructure:"GRPC_PORT"`
}

func LoadConfig(path string) {
	var config Config

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("Couldn't read config")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic("unable to unmarshal config")
	}

	App = &config

	//查看文件變化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changed")
	})
}

func GetConfigToString(key string) string {
	return reflect.ValueOf(App).Elem().FieldByName(key).String()
}

func GetConfigToInt(key string) int64 {
	return reflect.ValueOf(App).Elem().FieldByName(key).Int()
}
