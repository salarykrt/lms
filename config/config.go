package config

import (
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
)

// GlobalConfig is the global config variable
var GlobalConfig Config

// Config is the struct for the config file
type Config struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
	S3Bucket S3Config       `toml:"s3bucket"`
	Local    LocalConfig    `toml:"local"`
}

type DatabaseConfig struct {
	Server   string `toml:"server"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

type ServerConfig struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

type LocalConfig struct {
	BaseURL     string `toml:"base_url"`
	TempDocPath string `toml:"temp_doc_path"`
}

type S3Config struct {
	AccessKey    string `toml:"access_key"`
	SecretKey    string `toml:"secret_key"`
	Region       string `toml:"region"`
	Bucket       string `toml:"bucket"`
	StorageClass string `toml:"storage_class"`
}

func LoadConfig(file string) error {
	_, err := toml.DecodeFile(file, &GlobalConfig)
	return err
}

func GetEnvironment() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "development"
	}
	return env
}

func InitEnv() error {
	env := GetEnvironment()
	var configFileName string

	switch env {
	case "development":
		configFileName = "config/config_dev.toml"
	case "preprod":
		configFileName = "config/config_preprod.toml"
	case "production":
		configFileName = "config/config_prod.toml"
	default:
		slog.Error("Unsupported environment:", slog.String("env", env))
	}

	// Load the configuration
	err := LoadConfig(configFileName)
	return err
}
