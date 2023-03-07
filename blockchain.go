package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block // genesisBlock property represents the first block added to the blockchain.
	chain        []Block
	difficulty   int // the difficulty property defines the minimum effort miners have to undertake to mine a block and include it in the blockchain
}

func (b Block) calculateHash() string {
	// Convert the block’s data to JSON format
	data, _ := json.Marshal(b.data)
	// Concatenate the block’s previous hash, data, timestamp, and proof of work (PoW)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	// Hash the earlier concatenation with the SHA256 algorithm
	blockHash := sha256.Sum256([]byte(blockData))
	// Return the hashing result in base 16, with lowercase letters for A-F
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
		fmt.Println(b.pow, b.hash)
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	// Collects the details of a transaction
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	// Creates a new block with the transaction details
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	// Mines the new block with the previous block hash, current block data, and generated PoW
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		// Since hash is calculated based on the previous hash, data, timestamp, and PoW, if the current block’s hash
		// is not equal to the calculated hash, the data has been tampered with
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(1)

	// record transactions on the blockchain for Alice, Bob, and John
	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println(blockchain.isValid())
}
