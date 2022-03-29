package routes

import (
	"net/http"
	"strings"
)

type HandlerBlueprint interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type HTTPMethod string

type Route struct {
	Path    string
	Handler HandlerBlueprint
	Method  string
}

var routes = []string{
	"/",
	"/view",
	"/test",
}

var methods = []string{
	"get",
	"post",
	"put",
	"patch",
}

// IsValid Checks whether the supplied route path exists in our array of supported routes.
func (r Route) IsValid() bool {
	return IsRouteSupported(r.Path)
}

func IsRouteSupported(route string) bool {
	var supportedRoutes string
	for i, path := range routes {
		supportedRoutes += path
		if len(routes)-1 > i {
			supportedRoutes += "|"
		}
	}
	//var validPath = regexp.MustCompile("^/(" + supportedRoutes + ")/([a-zA-Z0-9]+)$")
	return strings.Contains(supportedRoutes, route)
}

func IsMethodSupported(method string) bool {
	var supportedMethods string
	for i, path := range methods {
		supportedMethods += path
		if len(methods)-1 > i {
			supportedMethods += "|"
		}
	}
	//var validPath = regexp.MustCompile("^/(" + supportedMethods + ")/([a-zA-Z0-9]+)$")
	return strings.Contains(supportedMethods, method)
}
