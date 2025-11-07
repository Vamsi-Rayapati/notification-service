package main

import (
	"github.com/smartbot/notification/api"
)

func main() {
	r := api.RegisterRoutes()
	r.Run(":4003")
}
