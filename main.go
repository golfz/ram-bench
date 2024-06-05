package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	dataSize = 1024 * 1024 * 1024 * 5 // 5 GB
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Allocate slices for benchmarking
	writeData := make([]byte, dataSize)
	readData := make([]byte, dataSize)
	copyData := make([]byte, dataSize)

	// Write benchmark
	start := time.Now()
	for i := range writeData {
		writeData[i] = byte(rand.Intn(256))
	}
	writeDuration := time.Since(start)
	fmt.Printf("Write Speed: Time taken to write %d GB of RAM: %v\n", dataSize/(1024*1024*1024), writeDuration)

	// Read benchmark
	start = time.Now()
	for i := range readData {
		readData[i] = writeData[i]
	}
	readDuration := time.Since(start)
	fmt.Printf("Read Speed: Time taken to read %d GB of RAM: %v\n", dataSize/(1024*1024*1024), readDuration)

	// Transfer benchmark
	start = time.Now()
	copy(copyData, writeData)
	transferDuration := time.Since(start)
	fmt.Printf("Transfer Rate: Time taken to copy %d GB of RAM: %v\n", dataSize/(1024*1024*1024), transferDuration)
}
