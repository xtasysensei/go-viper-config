package loadcfg

import (
	"fmt"

	"github.com/spf13/viper"
)

type PostgresqlConfig struct {
	Host string `mapstructure:"hostname"`
	Port string
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}
type DatabaseConfig struct {
	Psql PostgresqlConfig `mapstructure:"postgresql"`
}

type ServerConfig struct {
	Port string
}
type Config struct {
	Db  DatabaseConfig `mapstructure:"database"`
	SRV ServerConfig   `mapstructure:"server"`
}

func LoadTomlCfg() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("confi")
	viper.SetConfigType("toml")
	//viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s \n", err)
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	return c, nil
}
