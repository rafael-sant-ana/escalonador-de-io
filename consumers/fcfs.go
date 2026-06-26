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
		next, err := h.getNextFCFS()

		if err != nil {
			continue
		}

		distance := utils.Abs(next - h.currentPosition)

		h.totalMovement += distance
		h.currentPosition = next

		h.log("Total movement for FCFS:", h.totalMovement)
		time.Sleep(3 * time.Second)
	}
}

func (h *FCFSHandler) getNextFCFS() (int, error) {
	if len(h.requests) == 0 {
		return -1, errors.New("No requests to getNext from")
	}
	next := h.requests[0]

	h.requests = h.requests[1:]

	return next, nil
}

func (h *FCFSHandler) log(a ...any) {
	fmt.Print("[FCFSHandler] ")
	fmt.Println(a...)
}
