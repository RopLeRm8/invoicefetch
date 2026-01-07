package server

import (
	"fmt"
	"net/http"
	"sync"
)

var once sync.Once

func Init(codeCh chan string) {
	once.Do(func() {
		http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
			token := r.URL.Query().Get("code")
			codeCh <- token
			fmt.Fprintf(w, "Code received %s", token)
		})

		go func() {
			err := http.ListenAndServe(":8080", nil)
			if err != nil {
				panic(err)
			}
		}()
	})
}
