package core

import (
	"github.com/glebarez/sqlite"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

// NewDatabase creates a new database connection.
func NewDatabase(conf Config, log Logger) Database {
	var (
		dsn = conf.DB.DSN
		db  *gorm.DB
		err error
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
	return Database{DB: db}
}
