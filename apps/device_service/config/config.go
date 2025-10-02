package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"db"`

	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`

	RMQ struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"rmq"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Note: .env file not found, using environment variables only")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

func (c *Config) GetSQLxDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DB.Host,
		c.DB.Port,
		c.DB.User,
		c.DB.Password,
		c.DB.Name,
	)
}

func (c *Config) GetAMqpDSN() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		c.RMQ.User,
		c.RMQ.Password,
		c.RMQ.Host,
		c.RMQ.Port,
	)
}

func (c *Config) GetServerDSN() string {
	return fmt.Sprintf(
		"%s:%d",
		c.Server.Host,
		c.Server.Port,
	)
}

func setDefaults() {
	viper.SetDefault("db.host", "postgres")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.name", "device_service")
	viper.SetDefault("db.user", "postgres")
	viper.SetDefault("db.password", "postgres")

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.host", "0.0.0.0")

	viper.SetDefault("rmq.host", "rabbitmq")
	viper.SetDefault("rmq.port", 5672)
	viper.SetDefault("rmq.user", "guest")
	viper.SetDefault("rmq.password", "guest")
}
