package main

import (
	"fmt"
	"time"
)

type DiskInfo struct {
	maxDiskBytes    int
	initialPosition int
}

type IoHandler struct {
	diskInfo DiskInfo
	requests []int
}

type FCFSHandler struct {
	IoHandler
}

func NewFCFSHandler(diskInfo DiskInfo) *FCFSHandler {
	return &FCFSHandler{
		IoHandler: IoHandler{
			diskInfo: diskInfo,
			requests: []int{},
		},
	}
}

func (h *FCFSHandler) listenForMessages(requests chan int) {
	for request := range requests {

		fmt.Println("Got Request: ", request)

		h.requests = append(h.requests, request)
	}
}

func (h *FCFSHandler) handle() {
	for {
		totalMovementFCFS := 0
		fmt.Println("Total movement for FCFS:", totalMovementFCFS)

		time.Sleep(3 * time.Second)
	}
}

func getNextFCFS(requests []int) int {
	return requests[0]
}
