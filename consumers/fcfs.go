package consumers

import (
	"errors"
	"fmt"
	"io-scheduling/disk"
	"io-scheduling/utils"
	"time"
)

type FCFSHandler struct {
	IoHandler
	totalMovement int
}

func NewFCFSHandler(diskInfo disk.DiskInfo) *FCFSHandler {
	return &FCFSHandler{
		IoHandler: IoHandler{
			diskInfo:        diskInfo,
			requests:        []int{},
			currentPosition: 0,
		},
		totalMovement: 0,
	}
}

func (h *FCFSHandler) ListenForAccesses(requests chan int) {
	go func() {
		for request := range requests {
			h.log("Got access: ", request)

			h.requests = append(h.requests, request)
		}
	}()
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
