package scheduler

import (
	"fmt"
	"io-scheduling/disk"
	"io-scheduling/scheduler/strategies"
	"io-scheduling/utils"
	"time"
)

type Scheduler struct {
	diskInfo        disk.DiskInfo
	requests        []int
	currentPosition int

	ss            strategies.SchedulerStrategy
	totalMovement int
}

func NewScheduler(
	diskInfo disk.DiskInfo,
	initialPosition int,
	ss strategies.SchedulerStrategy,
) *Scheduler {
	return &Scheduler{
		diskInfo:        diskInfo,
		requests:        []int{},
		currentPosition: initialPosition,

		ss:            ss,
		totalMovement: 0,
	}
}

func (s *Scheduler) ListenForAccesses(requests chan int) {
	for request := range requests {
		s.requests = append(s.requests, request)
	}
}

func (s *Scheduler) Handle() {
	for {
		next, requests, err := s.ss.GetNext(s.requests, s.currentPosition)

		s.requests = requests

		if err != nil {
			continue // Isso causa gasto inutil de CPU
		}

		distance := utils.Abs(next - s.currentPosition)

		s.totalMovement += distance
		s.currentPosition = next

		s.logStatus()
		time.Sleep(3 * time.Second)
	}
}

func (s *Scheduler) logStatus() {
	fmt.Println(
		"[%s] Total Movement: %d | Current Position: %d | Requests List: %v",
		s.ss.Name(),
		s.totalMovement,
		s.currentPosition,
		s.requests,
	)
}
