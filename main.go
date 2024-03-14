package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"time"
)

const (
	ISO8601 = "2006-01-02T15:04:05Z"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		panic("PORT required")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]any{
			"date": time.Now().UTC().Format(ISO8601),
		})

		if err != nil {
			fmt.Println("Error encoding JSON: ", err)
		}
	})

	fmt.Printf("Listening on :%s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
