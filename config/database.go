package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// DatabaseConfiguration hold the values required to connect with the database
type DatabaseConfiguration struct {
	DbType      string `mapstructure:"DATABASE_TYPE"`
	DbUsername  string `mapstructure:"DATABASE_USERNAME"`
	DbPassword  string `mapstructure:"DATABASE_PASSWORD"`
	DbName      string `mapstructure:"DATABASE_NAME"`
	DbHost      string `mapstructure:"DATABASE_HOST"`
	DbPort      string `mapstructure:"DATABASE_PORT"`
	DatabaseURL string
}

//DatabaseConfig represents the database configuration
var DatabaseConfig DatabaseConfiguration

//GetDatabaseConfig represents the database configuration
func GetDatabaseConfig() (err error) {

	viper.SetConfigFile(`config.yml`)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&DatabaseConfig)
	if err != nil {
		return
	}

	DatabaseConfig.DatabaseURL = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", DatabaseConfig.DbUsername, DatabaseConfig.DbPassword, DatabaseConfig.DbHost, DatabaseConfig.DbPort, DatabaseConfig.DbName)

	return

}
