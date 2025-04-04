package main

import "github.com/miya8060/ripple-backend/internal/api"

func main() {
	r := api.SetupRouter()
	r.Run(":8080")
}
