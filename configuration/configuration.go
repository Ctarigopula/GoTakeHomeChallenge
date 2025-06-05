package configuration

import (
	"fmt"
	"os"

	gorm2 "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"take-home-challenge/coordinators"
)

type Config struct {
	DB                  *gorm.DB
	DB2                 *gorm2.DB
	MessagesCoordinator coordinators.MessagesCoordinator
	UsersCoordinator    coordinators.UsersCoordinator
}

func Load() (Config, error) {
	dbDSN := os.Getenv("DB_DSN")
	c := postgres.Config{
		DSN: dbDSN,
	}
	db, err := gorm.Open(postgres.New(c), nil)
	if err != nil {
		return Config{}, fmt.Errorf("failed to connect to database: %w", err)
	}
	db2, err := gorm2.Open("postgres", dbDSN+"?sslmode=disable")
	if err != nil {
		return Config{}, fmt.Errorf("failed to connect to database 2: %w", err)
	}
	conf := Config{
		DB:                  db,
		DB2:                 db2,
		MessagesCoordinator: coordinators.NewMessagesCoordinator(db2),
		UsersCoordinator:    coordinators.NewUsersCoordinator(db),
	}
	return conf, nil
}
