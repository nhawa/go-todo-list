package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-todo-list/application"
	"strings"
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetDefault("application.env", "development")
	viper.SetDefault("application.jwt.tokenTtl", 24)
	viper.SetDefault("application.jwt.refreshTtl", 24)
	viper.SetDefault("application.jwt.secret", "ottencofffee")

	applicationEnv := viper.Get("application.env")

	log.Info(fmt.Sprintf("Running on %s mode", applicationEnv))
	if nil == applicationEnv || "development" == applicationEnv {
		viper.SetConfigFile(`config.yml`)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

	}
}

func main() {
	apps := application.New()
	apps.ListenAndServe()
}
