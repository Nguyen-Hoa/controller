package main

import (
	manager "github.com/Nguyen-Hoa/manager"
)

func main() {
	g_manager, err := manager.NewManager("./config.json")
	if err != nil {
		return
	}
	g_manager.Start()
}
