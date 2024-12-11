package main

import "net/http"

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}))

	http.ListenAndServe(":8080", http.DefaultServeMux)
}
