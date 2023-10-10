package main

import (
	"fmt"
	"log"
	"net/http"
	"user-api/api/handler"
	"user-api/config"
	"user-api/middleware"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares in order to a http.HandlerFunc.
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func main() {
	// Config
	config.Load()
	host, port := config.C.ServerHost, config.C.ServerPort
	address := host + ":" + port

	// Middlewares
	commonMiddlewares := []Middleware{
		middleware.LoggingMiddleware,
		middleware.CORSMiddleware,
	}

	authMiddlewares := append(commonMiddlewares, middleware.JWTMiddleware)

	// Routes
	http.HandleFunc("/register", Chain(handler.RegisterUserHandler, commonMiddlewares...))
	http.HandleFunc("/profile", Chain(handler.ProfileHandler, authMiddlewares...))
	http.HandleFunc("/profile/update", Chain(handler.UpdateUserHandler, authMiddlewares...))
	http.HandleFunc("/profile/delete", Chain(handler.DeleteUserHandler, authMiddlewares...))
	http.HandleFunc("/login", Chain(handler.LoginHandler, commonMiddlewares...))
	http.HandleFunc("/logout", Chain(handler.LogoutHandler, authMiddlewares...))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Printf("Server is up and listening on port: %s\n", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
