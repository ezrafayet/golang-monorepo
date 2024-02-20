package main

import (
	"service-b/internal/httpserver"
)

func main() {
	if err := httpserver.Start(); err != nil {
		panic(err)
	}
}
