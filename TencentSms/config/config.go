package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	viper *viper.Viper
}

var Confs *config
var Secrets *config

func init() {
	Confs = &config{
		viper: getConf(),
	}
	Secrets = &config{
		viper: getSecret(),
	}
}

func getConf() *viper.Viper {
	v := viper.New()
	v.SetConfigName("confs")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/conf")
	err := v.ReadInConfig()
	fmt.Print(err)
	return v
}
func getSecret() *viper.Viper {
	v := viper.New()
	v.SetConfigName("secrets")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/secret")
	_ = v.ReadInConfig()
	return v
}
func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}
