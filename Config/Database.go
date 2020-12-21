package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DBConfig struct {
	DB *gorm.DB
}

type DBConfig struct {
	Host     string
	Port     uint
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {

}
