package example

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/visstars7/engine"
)

func TestExampleServer(t *testing.T) {
	router := engine.NewEngine()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Anya Forger.")
	})

	router.Use(SimpleMiddleware)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
