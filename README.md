# Blockchain Implementation in Go

This is a self-written blockchain implementation in Go that demonstrates the core concepts of blockchain technology, including:

- Block structure and hashing
- Proof of Work algorithm
- Blockchain persistence with BoltDB
- REST API for blockchain interactions

## Features

- **Block Structure**: Each block contains timestamp, data, previous block hash, and its own hash
- **Proof of Work**: Implements a basic Proof of Work algorithm for mining blocks
- **Persistence**: Uses BoltDB for persistent storage of the blockchain
- **REST API**: Provides HTTP endpoints for interacting with the blockchain
- **Command-line Interface**: Includes a command-line version for direct interaction

## Prerequisites

- Go 1.18 or higher
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd hw21
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Command-line Version

Run the command-line version to see the blockchain in action:

```bash
go run main.go
```

This will create a blockchain with a genesis block and add three sample blocks, then display all blocks.

### API Server

Run the API server to interact with the blockchain through HTTP requests:

```bash
go run main_api.go
```

The API server will start on `http://localhost:8080`.

#### API Endpoints

- `GET /blocks` - Get all blocks in the blockchain
- `POST /blocks` - Add a new block to the blockchain
- `GET /blocks/{hash}` - Get a specific block by its hash

#### Example API Usage

1. Get all blocks:
   ```bash
   curl -X GET http://localhost:8080/blocks
   ```

2. Add a new block:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"data":"Your data here"}' http://localhost:8080/blocks
   ```

3. Get a specific block:
   ```bash
   curl -X GET http://localhost:8080/blocks/{block_hash}
   ```

## Project Structure

```
.
├── main.go              # Command-line entry point
├── main_api.go          # API server entry point
├── api_server.go        # API server implementation
├── blockchain/          # Blockchain package
│   ├── block.go         # Block structure and methods
│   ├── blockchain.go    # Blockchain structure and methods
│   ├── database.go      # Database operations
│   ├── iterator.go      # Blockchain iterator
│   └── proofofwork.go   # Proof of Work implementation
├── go.mod               # Go module definition
└── go.sum               # Go module checksums
```

## How It Works

### Block Structure

Each block in the blockchain contains:
- `Timestamp`: When the block was created
- `Data`: The data stored in the block
- `PrevBlockHash`: Hash of the previous block
- `Hash`: Hash of the current block
- `Nonce`: Number used for Proof of Work

### Proof of Work

The Proof of Work algorithm requires mining a block by finding a hash that starts with a certain number of zeros. This process makes it computationally expensive to add blocks to the blockchain, which helps secure the network.

### Persistence

The blockchain is stored persistently using BoltDB, a pure Go key/value store. Each block is stored with its hash as the key, and the block data is serialized using Go's encoding/gob package.

## Development

### Building

To build the command-line version:
```bash
go build main.go
```

To build the API server:
```bash
go build main_api.go api_server.go
```

### Testing

Run the tests (if any):
```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.