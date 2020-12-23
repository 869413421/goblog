package bootstrap

import (
	"goblog/config"
	"goblog/pkg/model"
	"goblog/pkg/model/article"
	"goblog/pkg/model/user"
	"gorm.io/gorm"
)

func SetupDB() {
	dbConfig := config.LoadConfig()
	//建立连接池
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(dbConfig.Db.MaxConnections)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(dbConfig.Db.MaxIdeConnections)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(dbConfig.Db.ConnectionMaxLifeTime)

	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(&article.Article{}, &user.User{})
}
