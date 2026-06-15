package main

import (
	"fmt"
)

func main() {
	initialPosition := 256
	maxDiskBytes := 512
	requests := []int{315, 100, 15, 480, 0, 500, 15}

	processRequests := ProcessRequestsDTO{
		requests:        requests,
		initialPosition: initialPosition,
		maxDiskBytes:    maxDiskBytes,
	}

	totalMovementSCAN := scan(processRequests)
	totalMovementFCFS := fcfs(processRequests)
	totalMovementSSTF := sstf(processRequests)

	fmt.Println("Total movement for SCAN:", totalMovementSCAN)
	fmt.Println("Total movement for FCFS:", totalMovementFCFS)
	fmt.Println("Total movement for SSTF:", totalMovementSSTF)
}
