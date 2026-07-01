package consumers

import (
	"errors"
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
			h.requests = append(h.requests, request)
		}
	}()
}

func (h *SSTFHandler) Handle() {
	for {
		next, err := h.getNextSSTF()

		if err != nil {
			continue
		}

		distance := utils.Abs(next - h.currentPosition)

		h.totalMovement += distance
		h.currentPosition = next

		h.logStatus()
		time.Sleep(3 * time.Second)
	}
}

func (h *SSTFHandler) getNextSSTF() (int, error) {
	if len(h.requests) == 0 {
		return -1, errors.New("No requests to getNext from")
	}

	bestIdx := 0
	bestDistance := utils.Abs(h.currentPosition - h.requests[0])

	for idx, position := range h.requests {
		distance := utils.Abs(h.currentPosition - position)

		if distance < bestDistance {
			bestDistance = distance
			bestIdx = idx
		}
	}

	next := h.requests[bestIdx]

	h.requests = append(h.requests[:bestIdx], h.requests[bestIdx+1:]...) // pega a lista e desempacota porque espera append([]type, type, type, ...) = append([]type, ...type)

	return next, nil
}

func (h *SSTFHandler) logStatus() {
	h.log("Total movement for SSTF:", h.totalMovement, " | Current position", h.currentPosition, " | List of future accesses: ", h.requests)
}

func (h *SSTFHandler) log(a ...any) {
	fmt.Print("[SSTFHandler] ")
	fmt.Println(a...)
}
