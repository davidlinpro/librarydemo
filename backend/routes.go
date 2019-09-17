package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests(r *mux.Router) {
	r.Use(AddContentTypeMiddleware)

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")

	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users", getAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", editUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	r.HandleFunc("/books", createBookHandler).Methods("POST")
	r.HandleFunc("/books", getAllBooksHandler).Methods("GET")
	r.HandleFunc("/books/{id}", getBookHandler).Methods("GET")
	r.HandleFunc("/books/{id}", editBookHandler).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBookHandler).Methods("DELETE")

	r.HandleFunc("/inventory", createInventoryHandler).Methods("POST")
	r.HandleFunc("/inventory", getAllInventoryHandler).Methods("GET")
	r.HandleFunc("/inventory/{id}", getInventoryHandler).Methods("GET")
	r.HandleFunc("/inventory/{id}", editInventoryHandler).Methods("PUT")
	r.HandleFunc("/inventory/{id}/checkout/{guest_id}", checkoutInventoryHandler).Methods("PUT")
	r.HandleFunc("/inventory/{id}/return", returnInventoryHandler).Methods("PUT")
	r.HandleFunc("/inventory/{id}", deleteInventoryHandler).Methods("DELETE")
}

func AddContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
