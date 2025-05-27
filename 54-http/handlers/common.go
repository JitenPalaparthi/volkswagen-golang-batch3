package handlers

import "net/http"

func init() {
	println("handler calling --1 in utils")
}

func init() {
	println("Handler calling --2 in utils")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
	w.WriteHeader(http.StatusOK)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
	w.WriteHeader(http.StatusOK)
}
