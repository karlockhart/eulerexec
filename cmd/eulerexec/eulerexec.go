package main

import (
	"sync"

	"github.com/karlockhart/eulerexec/pkg/api"
	"github.com/karlockhart/eulerexec/pkg/config"
)

func main() {
	config.LoadConfig()
	var wg sync.WaitGroup
	wg.Add(1)
	go api.Start(&wg)
	wg.Wait()
}
