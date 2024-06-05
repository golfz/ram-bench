package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	loops    = 5
	dataSize = 1024 * 1024 * 1024 * 1 // 1 GB
)

func printBenchmark(name string, data []byte, fn func([]byte)) time.Duration {
	start := time.Now()

	fn(data)

	duration := time.Since(start)
	fmt.Printf("%s Time taken to operate on %d GB of RAM: %v\n", name, dataSize/(1024*1024*1024), duration)

	return duration
}

func writeBenchmark(data []byte) {
	for i := range data {
		data[i] = byte(rand.Intn(256))
	}
}

func readBenchmark(data []byte, origin []byte) {
	for i := range data {
		data[i] = origin[i]
	}
}

func copyBenchmark(data []byte, origin []byte) {
	copy(data, origin)
}

func main() {
	fmt.Println("Starting RAM benchmark...")

	var totalWriteTime, totalReadTime, totalTransferTime time.Duration

	for loop := 1; loop <= loops; loop++ {
		fmt.Printf("\n--- Loop %d ---\n", loop)

		writeData := make([]byte, dataSize)
		readData := make([]byte, dataSize)
		copyData := make([]byte, dataSize)

		totalWriteTime += printBenchmark("Write Speed:", writeData, writeBenchmark)
		totalReadTime += printBenchmark("Read Speed:", readData, func(data []byte) { readBenchmark(data, writeData) })
		totalTransferTime += printBenchmark("Transfer Rate:", copyData, func(data []byte) { copyBenchmark(data, writeData) })
	}

	// Print summary and calculate averages
	fmt.Println("\n--- Summary ---")
	fmt.Printf("Average Write Speed: %v\n", totalWriteTime/time.Duration(loops))
	fmt.Printf("Average Read Speed: %v\n", totalReadTime/time.Duration(loops))
	fmt.Printf("Average Transfer Rate: %v\n", totalTransferTime/time.Duration(loops))
}
