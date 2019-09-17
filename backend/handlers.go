package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	err := dbConn.Ping()
	if err != nil {
		io.WriteString(w, `{"alive": false}`)
	}
	io.WriteString(w, `{"alive": true}`)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&SimpleResponse{
		Data: "Backend API",
	})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearSession(w)
	json.NewEncoder(w).Encode(&OpResponse{
		Success: true,
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cardID := r.Form.Get("card_id")
	if cardID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	SetSession(w, cardID)
	id, _ := strconv.Atoi(cardID)
	SetLastLogin(id)
	json.NewEncoder(w).Encode(&OpResponse{
		Success: true,
	})
}
