package main

import (
	"fmt"
	"log"
	
	"github.com/nur1kmm/hw21/blockchain"
)

func main() {
	fmt.Println("Blockchain API Server")
	
	// Create a new blockchain
	bc := blockchain.NewBlockchain()
	defer bc.Close()
	
	// Create and start the API server
	apiServer := NewAPIServer(bc)
	
	fmt.Println("Blockchain API is running on http://localhost:8080")
	log.Fatal(apiServer.Start(":8080"))
}