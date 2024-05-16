package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"products/db"
	"products/web"
)

// Config represents the configuration structure
type Config struct {
	Name   string `json:"name"`
	Port   string `json:"port"`
	Host   string `json:"host"`
	Status string `json:"status"`
}

func main() {

	if err := db.InitDB(); err != nil {
		log.Fatal("Error loading Database", err)
	}
	defer db.Db.Close()

	config, err := readConfig("../config.json")
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}
	mux := web.StartServer()
	log.Printf("Server running on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, mux))

}

// readConfig reads and parses the configuration file

func readConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
