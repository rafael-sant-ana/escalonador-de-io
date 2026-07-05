package scheduler

import (
	"fmt"
	"io-scheduling/disk"
	"io-scheduling/scheduler/strategies"
	"io-scheduling/utils"
	"sync"
	"time"
)

type Scheduler struct {
	diskInfo        disk.DiskInfo
	requests        []int
	currentPosition int

	ss            strategies.SchedulerStrategy
	totalMovement int

	mu sync.Mutex
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
		s.mu.Lock()
		s.requests = append(s.requests, request)
		s.mu.Unlock()
	}
}

func (s *Scheduler) Handle() {
	fmt.Println("Activated handler for ", s.ss.Name())
	for {
		s.mu.Lock()
		next, requests, err := s.ss.GetNext(s.requests, s.currentPosition)

		s.requests = requests
		s.mu.Unlock()

		if err != nil {
			time.Sleep(100 * time.Millisecond)
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
	fmt.Printf(
		"[%s] Total Movement: %d | Current Position: %d | Requests List: %v \n",
		s.ss.Name(),
		s.totalMovement,
		s.currentPosition,
		s.requests,
	)
}
