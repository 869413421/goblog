package logger

import (
	. "goblog/config"
	"log"
	"os"
	"sync"
)

var Logger *log.Logger
var once sync.Once

func init() {
	once.Do(func() {
		config := LoadConfig()
		file, err := os.OpenFile(config.App.Log+"/blog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open log file ", err)
		}

		Logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

func Info(args ...interface{}) {
	Logger.SetPrefix("INFO ")
	Logger.Println(args...)
}

// 为什么不命名为 error？避免和 error 类型重名
func Danger(args ...interface{}) {
	Logger.SetPrefix("ERROR ")
	Logger.Println(args...)
}

func Warning(args ...interface{}) {
	Logger.SetPrefix("WARNING ")
	Logger.Println(args...)
}
