package main

import (
	"restApi/cmd/internal/app/config"
)

func main() {
	r := config.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
