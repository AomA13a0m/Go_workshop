package database

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

var (
	DBConn *gorm.DB
)
