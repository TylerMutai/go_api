package main

import (
	"api/api"
	"api/core/routes"
	"fmt"
	"os"
)

func main() {
	routes.BindToRoute(api.HandleIndex())
	err := routes.InitRoutes()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
