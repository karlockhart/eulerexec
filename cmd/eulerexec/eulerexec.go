package main

import (
	"sync"

	"github.com/coderzer0h/eulerexec/pkg/api"
	"github.com/coderzer0h/eulerexec/pkg/config"
)

func main() {
	config.LoadConfig()
	var wg sync.WaitGroup
	wg.Add(1)
	go api.Start(&wg)
	wg.Wait()
}
