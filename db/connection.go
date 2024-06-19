package db

import (
	"app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	if cfg.Driver != "mysql" {
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	return db, err
}
