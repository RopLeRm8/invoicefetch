package server

import (
	"fmt"
	"net/http"
)

func Init(codeCh chan string) {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("code")
		codeCh <- token
		fmt.Fprintf(w, "Code received %s", token)
	})
	http.ListenAndServe(":8080", nil)
}
