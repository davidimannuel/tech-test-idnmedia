package middlewares

import (
	"log"
	"net/http"
)

// Adapter is an alias so I dont have to type so much.
type Adapter func(http.Handler) http.Handler

// Adapt takes Handler funcs and chains them to the main handler.
func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	// The loop is reversed so the adapters/middleware gets executed in the same
	// order as provided in the array.
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}

// MethodLogger logs the method of the request.
func MethodLogger() Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			log.Printf("method=%s uri=%s\n", req.Method, req.RequestURI)
			next.ServeHTTP(res, req)
		})
	}
}
