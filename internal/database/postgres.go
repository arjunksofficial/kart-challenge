package database

import (
	"fmt"
	"os"

	"github.com/arjunksofficial/kart-challenge/internal/config"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PG *gorm.DB

func InitPostgres() error {
	dsn, err := ReadConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read Postgres config from environment variables")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to connect to Postgres database")
	}
	PG = db
	return nil
}

func GetPostgresDB() *gorm.DB {
	if PG == nil {
		err := InitPostgres()
		if err != nil {
			panic("failed to initialize Postgres database: " + err.Error())
		}
	}
	return PG
}
func ClosePostgres() error {
	if PG == nil {
		return nil // No connection to close
	}
	sqlDB, err := PG.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func ReadConfig() (string, error) {
	postgresHost := config.GetConfig().Postgres.Host
	postgresUser := config.GetConfig().Postgres.User
	postgresPassword := config.GetConfig().Postgres.Password
	postgresDB := config.GetConfig().Postgres.Database
	postgresPort := config.GetConfig().Postgres.Port

	fmt.Println("Postgres Config:", postgresHost, postgresUser, postgresPassword, postgresDB, postgresPort)
	if postgresHost == "" || postgresUser == "" || postgresPassword == "" || postgresDB == "" || postgresPort == "" {
		return "", os.ErrInvalid
	}
	dsn := "host=" + postgresHost + " user=" + postgresUser + " password=" + postgresPassword + " dbname=" + postgresDB + " port=" + postgresPort + " sslmode=disable"
	return dsn, nil
}
