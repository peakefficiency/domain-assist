package utils

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

// HandlerFunc represents a standard http handler function
type HandlerFunc func(http.ResponseWriter, *http.Request)

// RegisterHandlers automatically registers all handler functions from the api package
func RegisterHandlers(handlers ...HandlerFunc) {
	for _, handler := range handlers {
		// Get the function name using reflection
		funcName := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()

		// Extract just the handler name (e.g., "HelloHandler" from "package.HelloHandler")
		parts := strings.Split(funcName, ".")
		handlerName := parts[len(parts)-1]

		// Convert handler name to URL path
		// Remove "Handler" suffix and convert to lowercase
		path := "/api/" + strings.ToLower(strings.TrimSuffix(handlerName, "Handler"))

		// Register the handler
		http.HandleFunc(path, handler)
	}
}
