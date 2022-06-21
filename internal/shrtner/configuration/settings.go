package configuration

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Settings struct {
	Server   *Server
	Auth     *Auth
	Database *Database
	DNS      string
	Logging  *Logging
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

type Logging struct {
	Level   string
	Format  string
	Output  string
	DirPath string
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
			Host:     getStringEnvVarOrFile("server.host"),
			Port:     viper.GetInt("server.port"),
			Protocol: getStringEnvVarOrFile("server.protocol"),
		},
		Auth: &Auth{
			Enabled: getBoolEnvVarOrFile("auth.enabled"),
		},
		Database: &Database{
			Host:       getStringEnvVarOrFile("database.host"),
			Username:   getStringEnvVarOrFile("database.username"),
			Password:   getStringEnvVarOrFile("database.password"),
			DbName:     getStringEnvVarOrFile("database.db_name"),
			Enabled:    getBoolEnvVarOrFile("database.auth_enabled"),
			SkipSchema: getBoolEnvVarOrFile("database.skip_schema"),
		},
		DNS: getStringEnvVarOrFile("dns"),
		Logging: &Logging{
			Level:   getStringEnvVarOrFile("logging.level"),
			Format:  getStringEnvVarOrFile("logging.format"),
			Output:  getStringEnvVarOrFile("logging.output"),
			DirPath: getStringEnvVarOrFile("logging.directory"),
		},
	}
}

func getStringEnvVarOrFile(name string) string {
	s := os.Getenv(name)
	if s == "" {
		s = viper.GetString(name)
		if s == "" {
			panic(fmt.Sprintf("[WARNING] env variable %s does not exist.\n add configuration to environment variables or settings file.", name))
		}
	}

	return s
}

func getBoolEnvVarOrFile(name string) bool {
	s := os.Getenv(name)
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Printf(fmt.Sprintf("[WARNING] env variable %s does not exist. \nerror: %v. \nfallback: reading from settings.yaml...", name, err.Error()))
		v = viper.GetBool(name)
		return v
	}

	return v
}
