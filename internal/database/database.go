package database

import (
	"github.com/glebarez/sqlite"
	"github.com/xbmlz/gin-svelte-template/internal/config"
	"github.com/xbmlz/gin-svelte-template/internal/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a database connection.
type DB struct {
	ORM *gorm.DB
}

// NewDatabase creates a new database connection.
func NewDatabase(conf config.Config, log logger.Logger) DB {
	var (
		dsn = conf.DB.DSN
		err error
		db  *gorm.DB
	)
	switch conf.DB.Driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		log.Errorf("Unsupported database driver: %s", conf.DB.Driver)
	}
	if err != nil {
		log.Errorf("Failed to connect to database: %v", err)
	}

	return DB{
		ORM: db,
	}
}
