package consumers

import (
	"fmt"
	"time"
	"io-scheduling/disk"
)

type FCFSHandler struct {
	IoHandler
}

func NewFCFSHandler(diskInfo disk.DiskInfo) *FCFSHandler {
	return &FCFSHandler{
		IoHandler: IoHandler{
			diskInfo: diskInfo,
			requests: []int{},
		},
	}
}

func (h *FCFSHandler) ListenForMessages(requests chan int) {
	for request := range requests {

		fmt.Println("Got Request: ", request)

		h.requests = append(h.requests, request)
	}
}

func (h *FCFSHandler) Handle() {
	for {
		totalMovementFCFS := 0
		fmt.Println("Total movement for FCFS:", totalMovementFCFS)

		time.Sleep(3 * time.Second)
	}
}

func (h *FCFSHandler) getNextFCFS() int {

	return h.requests[0]
}
