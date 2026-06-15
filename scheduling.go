package main

import (
	"math"
	"slices"
)

type ProcessRequestsDTO struct {
	requests        []int
	maxDiskBytes    int
	initialPosition int
}

func scan(dto ProcessRequestsDTO) int {
	initialPosition, maxDiskBytes, requests := dto.initialPosition, dto.maxDiskBytes, dto.requests
	currentPosition := initialPosition
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

func fcfs(dto ProcessRequestsDTO) int {
	initialPosition, requests := dto.initialPosition, dto.requests
	currentPosition := initialPosition
	totalMovement := 0

	for _, request := range requests {
		totalMovement += abs(currentPosition - request)
		currentPosition = request
	}

	return totalMovement
}

func sstf(dto ProcessRequestsDTO) int {
	initialPosition, requests := dto.initialPosition, dto.requests
	totalMovement := 0
	currentPosition := initialPosition
	visited := make([]bool, len(requests))

	for range len(requests) {
		minDistance := math.MaxInt
		minDistanceIdx := -1

		for j := range len(requests) {
			if visited[j] {
				continue
			}
			distance := abs(currentPosition - requests[j])
			if distance < minDistance {
				minDistance = distance
				minDistanceIdx = j
			}
		}
		if minDistanceIdx != -1 {
			totalMovement += minDistance
			visited[minDistanceIdx] = true
			currentPosition = requests[minDistanceIdx]
		}
	}
	return totalMovement
}
