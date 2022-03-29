package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var boundToRoutes = make(map[string]HandlerBlueprint)

type Message struct {

	// [Name] Auto-injected when printing error
	Name string
	Body string

	// [Time] Auto-injected when printing both [error] and [message]
	Time int64
}

func BindToRoute(r Route) bool {
	if boundToRoutes[r.Path] != nil {
		fmt.Println("You cannot bind to an already bound route: " + r.Path + ", ignoring...")
		return false
	}
	if !IsMethodSupported(r.Method) {
		fmt.Println("The Method supplied: " + r.Method + ", is not supported. Ignoring.")
		return false
	}
	boundToRoutes[r.Path] = r.Handler
	return true
}

func InitRoutes() error {
	if len(boundToRoutes) < 1 {
		return errors.New("there are no bound routes. We'll now halt")
	}

	for k, v := range boundToRoutes {
		http.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			mainHandler(w, r, v)
		})
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}

func mainHandler(w http.ResponseWriter, r *http.Request, handler HandlerBlueprint) {
	// Set response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Check if path supplied matches our list of allowed routes
	if !IsRouteSupported(r.URL.Path) {
		PrintError(w, Message{Body: "Not Found"}, http.StatusNotFound)
		return
	}
	// TODO: Handle Auth.
	// TODO: Set up simple middleware (That can be registered with routes)

	// Check if the content-type header is 'application/json'
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		PrintError(w, Message{
			Body: "The supported content-type is application/json",
		}, http.StatusBadRequest)
		return
	}

	// invoke Handler with w and r
	handler.Handler(w, r)
}

func formatMessage(m Message) string {
	var jsonMessage string
	m.Time = time.Now().UnixNano()
	jsonMessageBytes, err := json.Marshal(m)
	if err != nil {
		jsonMessage = "{\"error\":\"" + string(err.Error()) + "\"}"
	} else {
		jsonMessage = string(jsonMessageBytes)
	}
	return jsonMessage
}

func PrintError(w http.ResponseWriter, m Message, status int) {
	// Both the 'name' and 'message' properties from variable 'm' are ignored.
	m.Name = "Error"
	w.WriteHeader(status)

	_, err := fmt.Fprintf(w, formatMessage(m))
	if err != nil {
		http.Error(w, formatMessage(m), http.StatusInternalServerError)
	}
}

func PrintMessage(w http.ResponseWriter, m Message) {
	m.Name = "Message"

	_, err := fmt.Fprintf(w, formatMessage(m))
	if err != nil {
		PrintError(w, m, http.StatusInternalServerError)
	}
}
