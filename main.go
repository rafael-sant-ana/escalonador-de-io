package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	maxDiskBytes := 512
	requests := []int{315, 100, 15, 480, 0, 500, 15}

	// SCAN
	totalMovementSCAN := scan(slices.Clone(requests), maxDiskBytes)

	// FCFS
	totalMovementFCFS := fcfs(slices.Clone(requests), maxDiskBytes)

	// SSTF
	totalMovementSSTF := sstf(slices.Clone(requests), maxDiskBytes)

	fmt.Println("Total movement for SCAN:", totalMovementSCAN)
	fmt.Println("Total movement for FCFS:", totalMovementFCFS)
	fmt.Println("Total movement for SSTF:", totalMovementSSTF)
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

	if len(left) > 0 && currentPosition < maxDiskBytes {
		totalMovement += abs(currentPosition - (maxDiskBytes - 1))
		currentPosition = maxDiskBytes - 1
	}

	slices.Reverse(left) // right->left

	for _, req := range left {
		totalMovement += abs(currentPosition - req)
		currentPosition = req
	}

	return totalMovement
}

func fcfs(requests []int, _maxDiskBytes int) int {
	currentPosition := 0
	totalMovement := 0

	for _, request := range requests {
		totalMovement += abs(currentPosition - request)
		currentPosition = request
	}

	return totalMovement
}

func sstf(requests []int, _maxDiskBytes int) int {
	totalMovement := 0
	position := 0
	visited := make([]bool, len(requests))

	for range len(requests) {
		minDistance := math.MaxInt
		minDistanceIdx := -1

		for j := range len(requests) {
			if visited[j] {
				continue
			}
			distance := abs(position - requests[j])
			if distance < minDistance {
				minDistance = distance
				minDistanceIdx = j
			}
		}
		if minDistanceIdx != -1 {
			totalMovement += minDistance
			visited[minDistanceIdx] = true
			position = requests[minDistanceIdx]
		}
	}
	return totalMovement
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
