package main

import (
	"os"
	"sync"

	eventService "github.com/devpablocristo/blankfactor/event-service/api"
)

const (
	defaultPort = "8080"
	port1       = "8081"
	port2       = "8082"
)

func main() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	reserveNumberPort := os.Getenv("number-manager_PORT")
	if reserveNumberPort == "" {
		reserveNumberPort = defaultPort
	}

	wg.Add(1)
	go eventService.StartApi(&wg, reserveNumberPort)
	wg.Wait()
}
