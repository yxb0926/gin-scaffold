package gormx

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

type Config struct {
	Debug       bool
	DBType      string
	DSN         string
	MaxLifeTime int
	MaxOpenCons int
	MaxIdleCons int
	TablePrefix string
}

func New(c *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch strings.ToLower(c.DBType) {
	case "mysql":
		dialector = mysql.Open(c.DSN)
	case "sqlite3":
		dialector = sqlite.Open(c.DSN)
	case "postgres":
		dialector = postgres.Open(c.DSN)
	default:
		return nil, errors.New("error db type, please check")
	}

	gConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gConfig)
	if err != nil {
		return nil, err
	}

	// 参数设置
	if c.Debug {
		db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifeTime) * time.Second)
	sqlDB.SetMaxIdleConns(c.MaxIdleCons)
	sqlDB.SetMaxOpenConns(c.MaxOpenCons)

	return db, nil
}
