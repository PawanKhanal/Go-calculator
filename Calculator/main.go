package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Handler to perform the calculation
func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON data
	var req struct {
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	if req.Expression == "" {
		http.Error(w, "Expression cannot be empty", http.StatusBadRequest)
		return
	}

	// Evaluate the expression
	result, err := evaluateExpression(req.Expression)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error evaluating expression: %v", err), http.StatusBadRequest)
		return
	}

	// Send the result as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"result": result,
	})
}

// Function to evaluate basic mathematical expressions
func evaluateExpression(expr string) (string, error) {
	// Check for supported operators
	if strings.Contains(expr, "+") {
		parts := strings.Split(expr, "+")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format")
		}
		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[0])
		}
		num2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[1])
		}
		return fmt.Sprintf("%f", num1+num2), nil
	}
	if strings.Contains(expr, "-") {
		parts := strings.Split(expr, "-")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format")
		}
		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[0])
		}
		num2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[1])
		}
		return fmt.Sprintf("%f", num1-num2), nil
	}
	if strings.Contains(expr, "*") {
		parts := strings.Split(expr, "*")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format")
		}
		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[0])
		}
		num2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[1])
		}
		return fmt.Sprintf("%f", num1*num2), nil
	}
	if strings.Contains(expr, "/") {
		parts := strings.Split(expr, "/")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format")
		}
		num1, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[0])
		}
		num2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return "", fmt.Errorf("invalid number: %v", parts[1])
		}
		if num2 == 0 {
			return "", fmt.Errorf("cannot divide by zero")
		}
		return fmt.Sprintf("%f", num1/num2), nil
	}

	return "", fmt.Errorf("unsupported operator")
}

// Main function to start the server
func main() {
	// Serve static files (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handler for /calculate
	http.HandleFunc("/calculate", calculateHandler)

	// Serve the main HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Start the server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
