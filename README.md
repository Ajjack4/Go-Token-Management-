# Token Management System

## Overview

This project is a simulation of a Token Management System written in Go. The system manages a pool of tokens and tracks their usage, resetting usage counts every 24 hours. It efficiently selects the least-used tokens for each operation and provides statistics on token usage.

## Features

- **Initialization**: A pool of 1000 tokens is created, each with a unique ID and a usage count initialized to zero.
- **Token Selection**: Tokens are selected based on the lowest usage count, with random selection among ties.
- **Usage Tracking**: Each selected token's usage count is incremented.
- **24-Hour Reset**: The usage counts are reset every 24 hours to ensure fair usage distribution.
- **Simulation**: The system allows simulating a specified number of operations and displays the results.

## Approach

### Data Structures

- **Token**: Represents a token with an ID, usage count, and index for heap operations.
- **TokenHeap**: A min-heap structure implemented using Go's `container/heap` package, maintaining tokens ordered by their usage count.
- **TokenPool**: Manages the collection of tokens, providing methods for token selection, usage update, and usage statistics display.

### Core Functions

- **NewTokenPool**: Initializes the token pool with 1000 tokens and sets up the min-heap.
- **SelectToken**: Selects the least-used token and updates its usage count.
- **checkAndReset**: Resets the usage counts if 24 hours have passed since the last reset.
- **SimulateOperations**: Simulates the specified number of token operations.
- **PrintStats**: Displays the usage statistics for all tokens, highlighting the least-used tokens.

## How to Run the Program

### Prerequisites

- Go installed on your machine (version 1.16 or higher recommended).

### Steps to Run

1. **Clone the Repository**

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Compile and Run**

   ```bash
   go run main.go
   ```

3. **Input Operations**

   When prompted, enter the number of operations to simulate. For example:

   ```
   Enter number of operations to simulate: 1000
   ```

4. **View Results**

   After the simulation, the program will display the usage statistics of all tokens and identify the least-used tokens.

### Example Output

```
Enter number of operations to simulate: 10
Starting simulation with 10 operations...

Token Usage Statistics:
Token 1: 2 uses
Token 2: 1 use
Token 3: 1 use
...
Token 1000: 0 uses

Least Used Token(s):
Token 2 (1 use)
Token 3 (1 use)
...
Token 1000 (0 uses)
```

## Future Improvements

- **Concurrency Handling**: Implementing concurrency to handle simultaneous token operations efficiently.
- **Persistence**: Adding persistent storage to maintain token states across program restarts.
- **API Interface**: Creating a RESTful API for external interaction with the token management system.

## Conclusion

This Token Management System is a simple yet efficient solution for managing and tracking token usage. The use of a min-heap ensures quick access to the least-used tokens, and the simulation feature allows for easy testing and verification of the system's functionality.

