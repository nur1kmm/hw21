package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Blockchain API Server")
	
	// Create a new blockchain
	bc := NewBlockchain()
	defer bc.Close()
	
	// Create and start the API server
	apiServer := NewAPIServer(bc)
	
	fmt.Println("Blockchain API is running on http://localhost:8080")
	log.Fatal(apiServer.Start(":8080"))
}