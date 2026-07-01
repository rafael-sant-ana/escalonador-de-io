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
	SSTFHandler := consumers.NewSSTFHandler(diskInfo)
	multiplexer := consumers.NewMultiplexerHandler(FCFSHandler, SSTFHandler)

	producer.ProduceRandomAccessesRequests(requests, diskInfo)

	multiplexer.ListenForAccesses(requests)

	go FCFSHandler.Handle()
	go SSTFHandler.Handle()

	select {}
}
