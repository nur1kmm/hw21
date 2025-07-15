package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nur1kmm/hw21/blockchain"
)

// APIServer represents the API server
type APIServer struct {
	blockchain *blockchain.Blockchain
	router     *mux.Router
}

// NewAPIServer creates a new API server
func NewAPIServer(blockchain *blockchain.Blockchain) *APIServer {
	router := mux.NewRouter()
	server := &APIServer{
		blockchain: blockchain,
		router:     router,
	}
	
	server.setupRoutes()
	
	return server
}

// setupRoutes sets up the API routes
func (s *APIServer) setupRoutes() {
	s.router.HandleFunc("/blocks", s.handleGetBlocks).Methods("GET")
	s.router.HandleFunc("/blocks", s.handleAddBlock).Methods("POST")
	s.router.HandleFunc("/blocks/{hash}", s.handleGetBlock).Methods("GET")
}

// handleGetBlocks handles GET /blocks request
func (s *APIServer) handleGetBlocks(w http.ResponseWriter, r *http.Request) {
	blocks := s.blockchain.GetBlocks()
	
	// Convert blocks to JSON-serializable format
	var response []map[string]interface{}
	for _, block := range blocks {
		blockData := map[string]interface{}{
			"hash":           fmt.Sprintf("%x", block.Hash),
			"prev_block_hash": fmt.Sprintf("%x", block.PrevBlockHash),
			"data":           string(block.Data),
			"timestamp":      block.Timestamp,
			"nonce":          block.Nonce,
		}
		response = append(response, blockData)
	}
	
	s.writeJSON(w, response)
}

// handleAddBlock handles POST /blocks request
func (s *APIServer) handleAddBlock(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	
	var requestData map[string]string
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	data, ok := requestData["data"]
	if !ok {
		http.Error(w, "Missing 'data' field", http.StatusBadRequest)
		return
	}
	
	s.blockchain.AddBlock(data)
	
	response := map[string]string{
		"message": "Block added successfully",
	}
	
	s.writeJSON(w, response)
}

// handleGetBlock handles GET /blocks/{hash} request
func (s *APIServer) handleGetBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	
	// For simplicity, we'll iterate through all blocks to find the one with matching hash
	// In a production system, you'd want to index blocks by hash for faster lookup
	bci := s.blockchain.Iterator()
	
	for {
		block := bci.Next()
		if block == nil {
			http.Error(w, "Block not found", http.StatusNotFound)
			return
		}
		
		if fmt.Sprintf("%x", block.Hash) == hash {
			response := map[string]interface{}{
				"hash":           fmt.Sprintf("%x", block.Hash),
				"prev_block_hash": fmt.Sprintf("%x", block.PrevBlockHash),
				"data":           string(block.Data),
				"timestamp":      block.Timestamp,
				"nonce":          block.Nonce,
			}
			
			s.writeJSON(w, response)
			return
		}
		
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	
	http.Error(w, "Block not found", http.StatusNotFound)
}

// writeJSON writes JSON response
func (s *APIServer) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// Start starts the API server
func (s *APIServer) Start(addr string) error {
	fmt.Printf("Starting API server on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}