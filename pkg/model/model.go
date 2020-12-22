package model

import (
	"github.com/go-sql-driver/mysql"
	. "goblog/config"
	"goblog/pkg/logger"
	"goblog/pkg/types"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gloger "gorm.io/gorm/logger"
)

type BaseModel struct {
	ID uint64
}

func (model BaseModel) GetStringID() string {
	return types.UInt64ToString(model.ID)
}

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

	DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: gloger.Default.LogMode(gloger.Info),
	})

	if err != nil {
		logger.Danger(err, "gorm open err")
	}
	return DB
}
