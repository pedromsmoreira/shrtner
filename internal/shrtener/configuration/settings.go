package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Settings struct {
	Server   *Server
	Auth     *Auth
	Database *Database
}

type Server struct {
	Host     string
	Port     int
	Protocol string
}

type Auth struct {
	Enabled bool
}

type Database struct {
	Host       string
	Username   string
	Password   string
	DbName     string
	Enabled    bool
	SkipSchema bool
}

func NewSettings(cfgFolder string) *Settings {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfgFolder)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error settings file: %w \n", err))
	}

	return &Settings{
		Server: &Server{
			Host:     viper.GetString("server.host"),
			Port:     viper.GetInt("server.port"),
			Protocol: viper.GetString("server.protocol"),
		},
		Auth: &Auth{
			Enabled: viper.GetBool("auth.enabled"),
		},
		Database: &Database{
			Host:       getStringOrDefault("database.host"),
			Username:   getStringOrDefault("database.username"),
			Password:   getStringOrDefault("database.password"),
			DbName:     getStringOrDefault("database.db_name"),
			Enabled:    viper.GetBool("database.auth_enabled"),
			SkipSchema: viper.GetBool("database.skip_schema"),
		},
	}
}

func getStringOrDefault(name string) string {
	s := viper.GetString(name)
	if s == "" {
		s = os.Getenv(name)
		if s == "" {
			panic(fmt.Sprintf("variable %s not set. add variable to environment variables or settings file.", name))
		}
	}

	return s
}
