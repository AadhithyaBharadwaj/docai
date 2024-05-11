package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FormData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Symptoms  string `json:"symptoms"`
}

func submitDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var formData FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Process the data
	fmt.Println("Received details:")
	fmt.Println("First Name:", formData.FirstName)
	fmt.Println("Last Name:", formData.LastName)
	fmt.Println("Age:", formData.Age)
	fmt.Println("Symptoms:", formData.Symptoms)

	// Return a simple response
	response := response(formData)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, response)
}

func main() {
	// Serve static files (HTML, CSS, JavaScript) from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("static")))

	// API endpoint for submitting details
	http.HandleFunc("/api/submit-details", submitDetailsHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func response(formData FormData) string {
	if formData.Symptoms == "fine" ||
		formData.Symptoms == "great" ||
		formData.Symptoms == "nice" ||
		formData.Symptoms == "okay" {
		return fmt.Sprint("Well, " + formData.FirstName + ", I'm glad to hear that you're feeling fine! If you have any concerns or questions in the future, feel free to reach out. Take care!\nThank you for using Doc AI. If you encounter any bugs or errors, please feel free to report to asdfghjk@yahoo.com so that we can squash them.\nIf you feel you need personal advice, please contact the below doctors:\nBabs: 34567894567")
	} else {
		return fmt.Sprint("Well, " + formData.FirstName + ", the solution is to take rest.")
	}
}
