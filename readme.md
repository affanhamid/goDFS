# goDFS - Distributed File Storage System

A peer-to-peer distributed file storage system built in Go, featuring encryption, content-addressable storage, and automatic replication across network nodes.

## Features

- **Peer-to-Peer Architecture**: Decentralized network where nodes can communicate with each other
- **End-to-End Encryption**: AES-CTR encryption for secure file storage and transmission
- **Content-Addressable Storage**: Files are stored using SHA-1 hash-based paths for deduplication
- **Automatic Replication**: Files are automatically replicated across network peers
- **TCP Transport Layer**: Custom P2P transport implementation with message streaming
- **Concurrent Operations**: Thread-safe operations with goroutines for handling multiple connections

## Architecture

The system consists of several key components:

- **FileServer**: Main orchestrator handling file operations and peer management
- **Store**: Content-addressable storage layer with configurable path transformation
- **P2P Transport**: TCP-based communication layer for peer discovery and messaging
- **Crypto**: AES-CTR encryption/decryption for secure file handling

## Technology Stack

- **Language**: Go 1.24.1
- **Encryption**: AES-CTR with random IV generation
- **Hashing**: SHA-1 for content addressing, MD5 for key hashing
- **Networking**: Custom TCP transport layer
- **Serialization**: Go's built-in `gob` encoding for message passing

## Getting Started

### Prerequisites
- Go 1.24.1 or later

### Building and Running

```bash
# Build the application
make build

# Run the distributed file system
make run

# Run tests
make test
```

### Example Usage

The system automatically starts three nodes:
- Node 1: `:3000` (bootstrap node)
- Node 2: `:3001` (bootstrap node) 
- Node 3: `:3002` (connects to bootstrap nodes)

Files are automatically stored, replicated, and retrieved across the network with encryption.

## Key Implementation Details

- **Path Transformation**: Files are stored using SHA-1 hash-based directory structures for efficient content addressing
- **Message Broadcasting**: Uses Go's `gob` encoding for efficient peer-to-peer communication
- **Stream Handling**: Custom stream management for large file transfers
- **Peer Management**: Thread-safe peer connection handling with automatic discovery
- **Encryption**: Each file is encrypted with a unique IV for security

## Project Structure

```
goDFS/
├── main.go              # Application entry point and node orchestration
├── server.go            # FileServer implementation with P2P logic
├── store.go             # Content-addressable storage layer
├── crypto.go            # Encryption/decryption utilities
├── p2p/                 # Peer-to-peer transport layer
│   ├── transport.go     # Transport interface definitions
│   ├── tcp_transport.go # TCP implementation
│   ├── message.go       # Message types and encoding
│   └── handshake.go     # Peer handshake protocol
└── bin/                 # Compiled binary output
```
