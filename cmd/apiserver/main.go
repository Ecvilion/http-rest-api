package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/http-rest-api/internal/app/apiserver"
)

/*
Хотим чтобы путь к файлу, который будет являться нашим конфигом (toml файл)

выступал в качестве флага при запуске нашего бинарника
*/
var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
