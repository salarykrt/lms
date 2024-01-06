package db

import (
	"fmt"
	configuration "salarykart/config"

	"github.com/gookit/config/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conn struct {
	db *gorm.DB
}

func NewConn() (*Conn, error) {
	// Access the database configuration
	dbConfig := configuration.GlobalConfig.Database

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.String("MYSQL_USER", dbConfig.Username),
			config.String("MYSQL_PASSWORD", dbConfig.Password),
			config.String("MYSQL_HOST", dbConfig.Server),
			config.Int("MYSQL_PORT", dbConfig.Port),
			config.String("MYSQL_FM_DB", dbConfig.DBName),
		),
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	return &Conn{db}, err
}

func (c *Conn) GetDB() *gorm.DB {
	return c.db
}
