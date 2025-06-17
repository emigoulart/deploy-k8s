package main

import "net/http"

func main() {
	// This is a placeholder for the main function.
	// You can add your application logic here.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
