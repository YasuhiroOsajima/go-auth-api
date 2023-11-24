package main

import "github.com/YasuhiroOsajima/go-auth-api/internal/infrastructure"

func main() {
	router := infrastructure.NewRouter()
	router.Run()
}
