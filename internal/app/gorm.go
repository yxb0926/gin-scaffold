package app

import (
	"errors"
	"gorm.io/gorm"
	"strings"
	"yxb.com/gin-scaffold/v2/internal/app/config"
	"yxb.com/gin-scaffold/v2/pkg/gormx"
)

// InitGormDB 根据不同的DBType创建不同类型的db实例
func InitGormDB() (*gorm.DB, error) {
	cfg := config.C
	var dsn string

	switch strings.ToLower(cfg.Gorm.DBType) {
	case "mysql":
		dsn = cfg.MySQL.DSN()
	case "sqlite3":
		dsn = cfg.Sqlite3.DSN()
	default:
		return nil, errors.New("unknown db type, please check")
	}

	return gormx.New(&gormx.Config{
		DSN:         dsn,
		DBType:      cfg.Gorm.DBType,
		Debug:       cfg.Gorm.Debug,
		MaxOpenCons: cfg.Gorm.MaxOpenConns,
		MaxIdleCons: cfg.Gorm.MaxIdleConns,
		MaxLifeTime: cfg.Gorm.MaxLifetime,
		TablePrefix: cfg.Gorm.TablePrefix,
	})
}
