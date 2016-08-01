package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Config struct {
	Listen string
}

func LoadConfig() Config {
	var configFile string
	flag.StringVar(&configFile, "c", "./config.json", "configure file")
	flag.Parse()

	f, err := os.Open(configFile)
	if err != nil {
		log.Fatalln("read configure error:", err)
	}

	var c Config
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		log.Fatalln("decode configure error:", err)
	}

	return c
}
