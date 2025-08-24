package database

import ()

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	_ "modernc.org/sqlite"  // This forces use of CGO-free driver
)

var(
	DBConn *gorm.DB
)

func ConnectDatabase() {
	d, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	DBConn = d
}

func GetDB() *gorm.DB {
	return DBConn
}
