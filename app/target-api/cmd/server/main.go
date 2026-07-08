package main

import (
	"fmt"
	"net/http"

	"github.com/gabriellim314/loadforge/app/target-api/internal/handlers"
)

func main() {

    http.HandleFunc("/health", handlers.Health)

	fmt.Println("🚀 Target API running on :8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}