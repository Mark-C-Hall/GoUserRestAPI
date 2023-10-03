package main

import (
	"fmt"
	"log"
	"net/http"
	"user-api/api/handler"
	"user-api/config"
	"user-api/middleware"
)

func main() {
	// Config
	config.Load()
	host, port := config.C.ServerHost, config.C.ServerPort
	address := host + ":" + port

	// Routes
	http.HandleFunc("/register", handler.RegisterUserHandler)
	http.HandleFunc("/profile", middleware.JWTMiddleware(handler.ProfileHandler))
	http.HandleFunc("/profile/update", middleware.JWTMiddleware(handler.UpdateUserHandler))
	http.HandleFunc("/profile/delete", middleware.JWTMiddleware(handler.DeleteUserHandler))
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/logout", middleware.JWTMiddleware(handler.LogoutHandler))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Printf("Server is up and listening on port: %s", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
