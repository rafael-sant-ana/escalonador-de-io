package strategies

import (
	"errors"
)

type FCFSStrategy struct{}

func NewFCFSStrategy() *FCFSStrategy {
	return &FCFSStrategy{}
}

func (s *FCFSStrategy) Name() string {
	return "FCFS Strategy"
}

func (s *FCFSStrategy) GetNext(requests []int, _currentPosition int) (int, []int, error) {
	if len(requests) == 0 {
		return -1, []int{}, errors.New("No requests to getNext from")
	}
	next := requests[0]

	requests = requests[1:]

	return next, requests, nil
}
