package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	result := Add(10, 5)

	fmt.Fprintf(
		w,
		"Go Jenkins Demo App - Running!<br>10 + 5 = %d",
		result,
	)
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on http://localhost:8085")

	err := http.ListenAndServe(":8085", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}