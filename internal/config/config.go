package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type RedisCfg struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type PostgresCfg struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Config struct {
	GinMode  string      `mapstructure:"gin_mode"`
	Port     string      `mapstructure:"port"`
	Env      string      `mapstructure:"env"`
	Redis    RedisCfg    `mapstructure:"redis"`
	Postgres PostgresCfg `mapstructure:"postgres"`
}

var appCfg *Config

func LoadConfig() error {
	// Load .env into OS environment
	_ = godotenv.Load()

	// Configure Viper to read from env
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // For nested structs

	// Explicit binds for nested structs
	bindEnvs := map[string]string{
		"gin_mode":          "GIN_MODE",
		"port":              "PORT",
		"env":               "ENV",
		"redis.host":        "REDIS_HOST",
		"redis.port":        "REDIS_PORT",
		"redis.password":    "REDIS_PASSWORD",
		"redis.db":          "REDIS_DB",
		"postgres.host":     "POSTGRES_HOST",
		"postgres.port":     "POSTGRES_PORT",
		"postgres.user":     "POSTGRES_USER",
		"postgres.password": "POSTGRES_PASSWORD",
		"postgres.database": "POSTGRES_DATABASE",
	}

	for key, env := range bindEnvs {
		if err := viper.BindEnv(key, env); err != nil {
			log.Fatalf("Error binding env var %s: %v", env, err)
		}
	}

	// Unmarshal into struct
	if err := viper.Unmarshal(&appCfg); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}
	return nil
}

func GetConfig() *Config {
	if appCfg == nil {
		if err := LoadConfig(); err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	}
	return appCfg
}

func GetRedisConfig() RedisCfg {
	cfg := GetConfig()
	return cfg.Redis
}

func (c *Config) IsReady() bool {
	// Implement any readiness checks here, e.g., checking Redis connection
	// For now, we assume the service is ready if the config is loaded
	return c != nil
}

// PromoImporterConfig holds the configuration for the promo importer
type PromoImporterConfig struct {
	FileSources [3]string
	RedisCfg    RedisCfg
}

var PromoImporterCfg *PromoImporterConfig

func LoadPromoImporterConfig() {
	_ = godotenv.Load()

	PromoImporterCfg = &PromoImporterConfig{
		FileSources: [3]string{
			os.Getenv("FILE1_SOURCE"),
			os.Getenv("FILE2_SOURCE"),
			os.Getenv("FILE3_SOURCE"),
		},
		RedisCfg: RedisCfg{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0, // Default Redis DB
		},
	}

	for i, f := range PromoImporterCfg.FileSources {
		if f == "" {
			log.Fatalf("FILE%d_SOURCE is missing in env", i+1)
		}
	}
}

func GetPromoImporterConfig() PromoImporterConfig {
	if PromoImporterCfg == nil {
		LoadPromoImporterConfig()
	}
	return *PromoImporterCfg
}
