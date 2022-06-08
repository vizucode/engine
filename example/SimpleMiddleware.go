package example

import (
	"fmt"
	"net/http"
)

//Example make some middleware.
func SimpleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before Requested....")
		next.ServeHTTP(w, r)
		fmt.Println("After Requested....")
	})
}
