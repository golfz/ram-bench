package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	dataSize = 1024 * 1024 * 1024 * 1 // 1 GB
	loops    = 3
)

func main() {
	fmt.Println("Starting RAM benchmark...")

	var totalWriteDuration, totalReadDuration, totalTransferDuration time.Duration

	for loop := 1; loop <= loops; loop++ {
		fmt.Printf("\n--- Loop %d ---\n", loop)

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
		totalWriteDuration += writeDuration

		// Read benchmark
		start = time.Now()
		for i := range readData {
			readData[i] = writeData[i]
		}
		readDuration := time.Since(start)
		fmt.Printf("Read Speed: Time taken to read %d GB of RAM: %v\n", dataSize/(1024*1024*1024), readDuration)
		totalReadDuration += readDuration

		// Transfer benchmark
		start = time.Now()
		copy(copyData, writeData)
		transferDuration := time.Since(start)
		fmt.Printf("Transfer Rate: Time taken to copy %d GB of RAM: %v\n", dataSize/(1024*1024*1024), transferDuration)
		totalTransferDuration += transferDuration
	}

	// Print summary and calculate averages
	fmt.Println("\n--- Summary ---")
	fmt.Printf("Average Write Speed: %v\n", totalWriteDuration/time.Duration(loops))
	fmt.Printf("Average Read Speed: %v\n", totalReadDuration/time.Duration(loops))
	fmt.Printf("Average Transfer Rate: %v\n", totalTransferDuration/time.Duration(loops))
}
