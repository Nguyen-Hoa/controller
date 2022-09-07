package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	manager "github.com/Nguyen-Hoa/manager"
)

var g_manager = manager.Manager{}

func initManager() {
	// Open JSON config
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Failed to parse configuration file.")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config manager.ManagerConfig
	json.Unmarshal([]byte(byteValue), &config)
	if err = g_manager.Init(config); err != nil {
		log.Println("Failed to initialize manager!")
		log.Fatal(err)
	}

	log.Println("Manager initialized")
}

func main() {
	initManager()
	g_manager.Start()
}
