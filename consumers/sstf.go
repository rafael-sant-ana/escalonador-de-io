package consumers

import (
	"fmt"
	"io-scheduling/disk"
	"io-scheduling/utils"
	"time"
)

type SSTFHandler struct {
	IoHandler
	totalMovement int
}

func NewSSTFHandler(diskInfo disk.DiskInfo) *SSTFHandler {
	return &SSTFHandler{
		IoHandler: IoHandler{
			diskInfo:        diskInfo,
			requests:        []int{},
			currentPosition: 0,
		},
		totalMovement: 0,
	}
}

func (h *SSTFHandler) ListenForAccesses(requests chan int) {
	go func() {
		for request := range requests {
			h.log("Got access: ", request)
			h.requests = append(h.requests, request)
		}
	}()
}

func (h *SSTFHandler) Handle() {
	go func() {
		for {
			next, err := h.getNextSSTF()

			if err != nil {
				continue
			}

			distance := utils.Abs(next - h.currentPosition)

			h.totalMovement += distance
			h.currentPosition = next

			h.log("Total movement for SSTF:", h.totalMovement, " | List of future accesses: ", h.IoHandler.requests)
			time.Sleep(3 * time.Second)
		}
	}()
}

func (h *SSTFHandler) getNextSSTF() (int, error) {
	return h.IoHandler.requests[0], nil
}

func (h *SSTFHandler) log(a ...any) {
	fmt.Print("[SSTFHandler] ")
	fmt.Println(a...)
}
