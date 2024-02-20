package main

import (
	"service-a/internal/httpserver"
)

func main() {
	if err := httpserver.Start(); err != nil {
		panic(err)
	}
}
