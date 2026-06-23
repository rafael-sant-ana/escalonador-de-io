package main

import (
	"io-scheduling/consumers"
	"io-scheduling/producer"
)

func main() {
	initialPosition := 256
	maxDiskBytes := 512

	requests := make(chan int)

	diskInfo := consumers.NewDiskInfo(
		initialPosition,
		maxDiskBytes,
	)

	FCFSHandler := consumers.NewFCFSHandler(diskInfo)

	producer.ProduceRandomAccessesRequests(requests)

	FCFSHandler.ListenForMessages(requests)

	//FCFSHandler.Handle()
}
