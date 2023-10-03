package main

import (
	"fmt"
	"log"
	"net/http"
	"user-api/api/handler"
	"user-api/config"
)

func main() {
	// Config
	config.Load()
	host, port := config.C.ServerHost, config.C.ServerPort
	address := host + ":" + port

	// Routes
	http.HandleFunc("/register", handler.RegisterUserHandler)
	http.HandleFunc("/profile", handler.ProfileHandler)
	http.HandleFunc("/profile/update", handler.UpdateUserHandler)
	http.HandleFunc("/profile/delete", handler.DeleteUserHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/logout", handler.LogoutHandler)
	http.HandleFunc("/reset", handler.ResetHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Printf("Server is up and listening on port: %s", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
