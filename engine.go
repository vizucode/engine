package engine

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Struct that represent engine object.
type engine struct {
	httprouter.Router
	middlewares []func(http.Handler) http.Handler
}

//Initializing an engine object.
func NewEngine() *engine {
	return new(engine)
}

//Use method is for registering an middlewares
func (e *engine) Use(next func(http.Handler) http.Handler) {
	e.middlewares = append(e.middlewares, next)
}

//Method to implements contract with http.Handler
func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var clientRequest http.Handler = &e.Router

	for _, next := range e.middlewares {
		clientRequest = next(clientRequest)
	}

	clientRequest.ServeHTTP(w, r)
}
