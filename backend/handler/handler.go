package handler

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//checks server running
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Hello, Postman!"))
}
