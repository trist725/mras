package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}
}
