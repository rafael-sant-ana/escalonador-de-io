package main

func main() {
	initialPosition := 256
	maxDiskBytes := 512

	requests := make(chan int)

	diskInfo := DiskInfo{
		initialPosition: initialPosition,
		maxDiskBytes:    maxDiskBytes,
	}

	go func() {
		requests <- 100
	}()
	FCFSHandler := NewFCFSHandler(diskInfo)

	FCFSHandler.listenForMessages(requests)

	//FCFSHandler.handle()
}
