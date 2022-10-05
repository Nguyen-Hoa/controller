package main

import (
	manager "github.com/Nguyen-Hoa/manager"
)

func main() {

	experiments := []string{"./config.json", "./config.json"}
	for _, e := range experiments {
		g_manager, err := manager.NewManager(e)
		if err != nil {
			return
		}
		g_manager.Start()
		g_manager = manager.Manager{}
	}
}
