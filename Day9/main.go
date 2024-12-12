package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	secretMessage := os.Getenv("SECRET_MESSAGE")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!\nSecret message: %s", secretMessage)
	})
	http.HandleFunc("/andy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Andy!") // this will be printed as a response when you request /
	})
	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(":" + port, nil) 
}
