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
	Database int    `mapstructure:"database"`
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

	_ = godotenv.Load()
	// 2. Setup viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 3. Set defaults (optional)
	viper.SetDefault("port", "9000")
	viper.SetDefault("env", "dev")

	// 4. Read config.yaml if present
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config.yaml found, relying on env vars")
	}
	appCfg = &Config{}

	// 5. Unmarshal into struct
	if err := viper.Unmarshal(appCfg); err != nil {
		return err
	}

	return nil
}

// PromoImportedConfig

type PromoImporterConfig struct {
	FileSources [3]string
	RedisCfg    RedisCfg
}

var PromoImporterCfg PromoImporterConfig

func LoadPromoImporterConfig() {
	_ = godotenv.Load()

	PromoImporterCfg = PromoImporterConfig{
		FileSources: [3]string{
			os.Getenv("FILE1_SOURCE"),
			os.Getenv("FILE2_SOURCE"),
			os.Getenv("FILE3_SOURCE"),
		},
	}

	for i, f := range PromoImporterCfg.FileSources {
		if f == "" {
			log.Fatalf("FILE%d_SOURCE is missing in env", i+1)
		}
	}
}
