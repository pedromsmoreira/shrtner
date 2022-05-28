package configuration

import (
	"fmt"
	"github.com/spf13/viper"
)

type Settings struct {
	Server   Server
	Auth     Auth
	Database Database
}

type Server struct {
	Host string
	Port int
}

type Auth struct {
	Enabled bool
}

type Database struct {
	Host     string
	Username string
	Password string
	DbName   string
	Enabled  bool
}

func NewSettings() *Settings {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/shrtener/configuration")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error settings file: %w \n", err))
	}

	return &Settings{
		Server: Server{
			Port: viper.GetInt("server.port"),
		},
		Auth: Auth{
			Enabled: viper.GetBool("auth.enabled"),
		},
		Database: Database{
			Host:     viper.GetString("database.host"),
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			DbName:   viper.GetString("database.db_name"),
			Enabled:  viper.GetBool("database.auth_enabled"),
		},
	}
}
