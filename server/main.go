package main

import (
	"fmt"
	"net/http"
	"server/api"
)

func main() {
	fmt.Println("Hello, World!")

	srv := api.NewServer()
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		return
	}
}
