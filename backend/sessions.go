package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func GetToken() string {
	timestamp := time.Now().Unix()
	hash := md5.New()
	io.WriteString(hash, strconv.FormatInt(timestamp, 10))
	token := fmt.Sprintf("%x", hash.Sum(nil))
	return token
}

func GetID(request *http.Request) int {
	cookie, err := request.Cookie("session")
	if err != nil {
		log.Println("missing session cookie")
		return -1
	}
	cookieValue := make(map[string]string)
	err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
	if err != nil {
		log.Println("invalid session cookie")
		return -1
	}
	if cookieValue["id"] == "" {
		log.Println("no id in session cookie")
		return -1
	}
	result, _ := strconv.Atoi(cookieValue["id"])
	return result
}

func SetSession(response http.ResponseWriter, id string) {
	if id == "" {
		return
	}
	value := map[string]string{
		"id": id,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
