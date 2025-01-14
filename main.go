package main

import (
	"container/heap"
	"fmt"

	"sort"
	"time"
)

type Token struct {
	ID    int
	Usage int
	Index int
}
type TokenHeap []*Token

func (h TokenHeap) Len() int           { return len(h) }
func (h TokenHeap) Less(i, j int) bool { return h[i].Usage < h[j].Usage }
func (h TokenHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}
func (h *TokenHeap) Push(x interface{}) {
	n := len(*h)
	token := x.(*Token)
	token.Index = n
	*h = append(*h, token)
}

func (h *TokenHeap) Pop() interface{} {
	old := *h
	n := len(old)
	token := old[n-1]
	token.Index = -1
	*h = old[0 : n-1]
	return token
}

// TokenPool manages the collection of tokens
type TokenPool struct {
	tokens    TokenHeap
	lastReset time.Time
}

// NewTokenPool initializes a new pool with 1000 tokens
func NewTokenPool() *TokenPool {
	tokens := make(TokenHeap, 1000)
	for i := range tokens {
		tokens[i] = &Token{ID: i + 1, Usage: 0}
	}
	heap.Init(&tokens)
	return &TokenPool{
		tokens:    tokens,
		lastReset: time.Now(),
	}
}

// SelectToken chooses the least used token, randomly selecting among ties
func (tp *TokenPool) SelectToken() *Token {
	tp.checkAndReset()
	token := heap.Pop(&tp.tokens).(*Token)
	token.Usage++
	heap.Push(&tp.tokens, token)
	return token
}

// checkAndReset resets usage counts if 24 hours have passed
func (tp *TokenPool) checkAndReset() {
	if time.Since(tp.lastReset) >= 24*time.Hour {
		for i := range tp.tokens {
			tp.tokens[i].Usage = 0
		}
		heap.Init(&tp.tokens)
		tp.lastReset = time.Now()
	}
}

// GetLeastUsedTokens returns all tokens with the minimum usage count
func (tp *TokenPool) GetLeastUsedTokens() []*Token {
	minUsage := tp.tokens[0].Usage
	for _, t := range tp.tokens {
		if t.Usage < minUsage {
			minUsage = t.Usage
		}
	}

	var leastUsed []*Token
	for _, t := range tp.tokens {
		if t.Usage == minUsage {
			leastUsed = append(leastUsed, t)
		}
	}
	return leastUsed
}

// PrintStats displays usage statistics for all tokens
func (tp *TokenPool) PrintStats() {
	// Sort tokens by ID for consistent display
	sort.Slice(tp.tokens, func(i, j int) bool {
		return tp.tokens[i].ID < tp.tokens[j].ID
	})

	fmt.Println("\nToken Usage Statistics:")
	for _, t := range tp.tokens {
		fmt.Printf("Token %d: %d uses\n", t.ID, t.Usage)
	}

	leastUsed := tp.GetLeastUsedTokens()
	fmt.Println("\nLeast Used Token(s):")
	for _, t := range leastUsed {
		fmt.Printf("Token %d (%d uses)\n", t.ID, t.Usage)
	}
}

// SimulateOperations runs the specified number of token operations
func (tp *TokenPool) SimulateOperations(count int) {
	fmt.Printf("Starting simulation with %d operations...\n", count)

	for i := 0; i < count; i++ {
		token := tp.SelectToken()
		token.Usage++
	}
}

func main() {

	// Initialize token pool
	pool := NewTokenPool()

	// // Get number of operations from user
	var operations int
	fmt.Print("Enter number of operations to simulate: ")
	fmt.Scan(&operations)

	// // Run simulation
	pool.SimulateOperations(operations)

	// // Display results
	pool.PrintStats()

}
