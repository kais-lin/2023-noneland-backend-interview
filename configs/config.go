package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port string
	Mode string
}

func init() {
	err := os.Setenv("TZ", "UTC")
	if err != nil {
		panic(fmt.Errorf("fatal error configs file: set time zone to utc: %w", err))
	}

	viper.AutomaticEnv()
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath("./configs")

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("讀取設定檔出現錯誤，原因為：%v", err))
	}
}

func NewConfig() *Config {
	return &Config{
		Port: viper.GetString("port"),
		Mode: viper.GetString("mode"),
	}
}
