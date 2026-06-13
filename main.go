package main

import (
	"fmt"
	"slices"
)

func main() {
	maxDiskBytes := 512
	requests := []int {315, 100, 15, 480, 0, 500}

	// SCAN
	totalMovementSCAN := scan(requests, maxDiskBytes)

	// FCFS
	// totalMovementFCFS := fcfs(requests, maxDiskBytes)

	fmt.Println("Total movement for SCAN:", totalMovementSCAN)
	// fmt.Println("Total movement for FCFS:", totalMovementFCFS)
}

func scan(requests []int, maxDiskBytes int) int {
	currentPosition := 0
	totalMovement := 0
	slices.Sort(requests)

	// esquerda -> direita
	// direita -> esquerda
	// volta
	var left []int
	var right []int

	for _, req := range requests {
		if req < currentPosition {
			left = append(left, req)
		} else {
			right = append(right, req)
		}
	}

	for _, req := range right {
		totalMovement += abs(currentPosition - req)
		currentPosition = req
	}

	if len(left > 0) && currentPosition < maxDiskBytes {
		totalMovement += abs(currentPosition - (maxDiskBytes-1))
		currentPosition = maskDiskBytes-1
	}

	slices.Reverse(left) // right->left

	for _, req := range left {
		totalMovement += abs(currentPosition - req)
		currentPosition = req
	}


	return totalMovement
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

