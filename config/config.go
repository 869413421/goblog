package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type App struct {
	Address string
	Static  string
	Logs    string
}

type Configuration struct {
	App App
}

var config *Configuration
var once sync.Once

func LoadConfig() *Configuration {
	//通过单例加载配置文件
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatal("open config error", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatal("Decode Config Error ", err)
		}
	})

	return config
}
