package main

import (
	"fmt"
	"os"
)

func main() {
	var apiEndpoint string

	// Check if API Endpoint is provided as a command line argument
	if len(os.Args) > 1 {
		apiEndpoint = os.Args[1]
	} else {
		// Check if API Endpoint is provided as an environment variable
		apiEndpoint = os.Getenv("OLLAMA_API_ENDPOINT")
	}

	// Print error message if API Endpoint is not provided
	if apiEndpoint == "" {
		fmt.Println("Error: API Endpoint is not provided. Please provide it as a command line argument or set the OLLAMA_API_ENDPOINT environment variable.")
		return
	}

	// Print the API Endpoint if it is provided
	fmt.Println("API Endpoint:", apiEndpoint)
}
