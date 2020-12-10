package model

import (
	"github.com/go-sql-driver/mysql"
	. "goblog/config"
	"goblog/pkg/logger"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	config := LoadConfig()

	dbConfig := mysql.Config{
		Addr:                 config.Db.Address,
		User:                 config.Db.User,
		Passwd:               config.Db.Password,
		DBName:               config.Db.Database,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	dsn := dbConfig.FormatDSN()

	dialector := mysql2.New(mysql2.Config{
		DSN: dsn,
	})

	DB, err = gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		logger.Danger(err, "gorm open err")
	}
	return DB
}
