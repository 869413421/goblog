package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	. "goblog/config"
	"goblog/pkg/logger"
	"log"
	"sync"
)

var DB *sql.DB
var err error
var once sync.Once

func init() {
	once.Do(func() {
		config := LoadConfig()
		dbConfig := mysql.Config{
			Addr:                 config.Db.Address,
			User:                 config.Db.User,
			Passwd:               config.Db.Password,
			DBName:               config.Db.Database,
			Net:                  "tcp",
			AllowNativePasswords: true,
		}

		DB, err = sql.Open(config.Db.Driver, dbConfig.FormatDSN())
		if err != nil {
			log.Fatalln(err, "open mysql err")
		}
		//设置最大连接数
		DB.SetMaxOpenConns(config.Db.MaxConnections)
		//设置最大空闲连接数
		DB.SetMaxIdleConns(config.Db.MaxIdeConnections)
		//设置链接超时时间
		DB.SetConnMaxLifetime(config.Db.ConnectionMaxLifeTime)

		err = DB.Ping()
		if err != nil {
			log.Fatalln(err, "ping mysql err")
		}
	})
}

func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
); `

	_, err := DB.Exec(createArticlesSQL)
	if err != nil {
		logger.Danger(err, "create articles table err")
	}
}
