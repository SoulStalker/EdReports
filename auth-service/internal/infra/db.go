package infra

import (
	"github.com/SoulStalker/edreports-auth-servive/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DSN string
}

func InitDB(cfg DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN))
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Role{},
		&domain.Permission{},
		&domain.RolePermission{},
		&domain.UserRole{},
		&domain.UserStore{},
		&domain.UserCategory{},
		&domain.AuthKey{},
		&domain.RefreshToken{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
