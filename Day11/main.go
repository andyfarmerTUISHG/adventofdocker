package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}

func main() {
	port := os.Getenv("PORT")
	secretMessage := os.Getenv("SECRET_MESSAGE")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Record new visit
		timestamp := time.Now().Format(time.RFC3339) + "\n"
		os.MkdirAll("data", 0755)
		f, _ := os.OpenFile("data/visits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.WriteString(timestamp)
		f.Close()

		// Read and display all visits
		data, err := os.ReadFile("data/visits.txt")
		if err != nil {
			fmt.Fprintf(w, "Error reading visits!")
			return
		}
		// fmt.Fprintf(w, "Hello World!\nSecret message: %s", secretMessage)
		fmt.Fprintf(w, "Testing.... Hello World!\nSecret message: %s \n\nNew visit recorded at %s\n\nAll visits:\n%s", secretMessage, timestamp, string(data))
	})
	http.HandleFunc("/andy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Andy!") // this will be printed as a response when you request /
	})
	fmt.Printf("Listening on port %s\n", port)
	http.ListenAndServe(":" + port, nil) 
}
