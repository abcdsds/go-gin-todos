package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type dbConfig struct {
	host     string
	port     int
	user     string
	dbName   string
	password string
}

func (c dbConfig) URL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.user,
		c.password,
		c.host,
		c.port,
		c.dbName)
}

func NewDBConfig() *dbConfig {
	return &dbConfig{
		host:     "0.0.0.0",
		port:     3306,
		user:     "root",
		dbName:   "todos",
		password: "pass",
	}
}
