package main

import (
	"io-scheduling/consumers"
	"io-scheduling/disk"
	"io-scheduling/producer"
)

func main() {
	requests := make(chan int)

	initialPosition := 256
	maxDiskBytes := 512
	diskInfo := disk.NewDiskInfo(
		initialPosition,
		maxDiskBytes,
	)

	FCFSHandler := consumers.NewFCFSHandler(diskInfo)

	producer.ProduceRandomAccessesRequests(requests, diskInfo)

	FCFSHandler.ListenForAccesses(requests)
	FCFSHandler.Handle()
}
