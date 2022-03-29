package api

import (
	"api/core/routes"
	"net/http"
)

/*
Handler for the '/' URLs or index page
*/

const path = "/"

const method = "get"

type HandlerImplementation struct{}

func HandleIndex() routes.Route {
	return routes.Route{Path: path,
		Method:  method,
		Handler: HandlerImplementation{},
	}
}

func (_ HandlerImplementation) Handler(w http.ResponseWriter, r *http.Request) {
	routes.PrintMessage(w, routes.Message{Body: "We just hit index page! Viola!!"})
}
