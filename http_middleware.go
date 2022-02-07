package gof_go

import (
	"fmt"
	"net/http"
)

type HTTPMiddleware func(w http.ResponseWriter, r *http.Request, next http.Handler)

func LoggingMiddleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	fmt.Printf("REQUEST %v\n", r.URL.String())
	next.ServeHTTP(w, r)
}

func OkHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		fmt.Printf("failed to write to body: %v", err)
	}
}

func MidddlewareToHandler(middleware HTTPMiddleware, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		middleware(writer, request, next)
	})
}
