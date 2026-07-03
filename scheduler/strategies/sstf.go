package strategies

import (
	"errors"
	"io-scheduling/utils"
)

type SSTFStrategy struct{}

func NewSSTFStrategy() *SSTFStrategy {
	return &SSTFStrategy{}
}

func (s *SSTFStrategy) Name() string {
	return "SSTF Strategy"
}

func (h *SSTFStrategy) GetNext(requests []int, currentPosition int) (int, []int, error) {
	if len(requests) == 0 {
		return -1, []int{}, errors.New("No requests to getNext from")
	}

	bestIdx := 0
	bestDistance := utils.Abs(currentPosition - requests[0])

	for idx, position := range requests {
		distance := utils.Abs(currentPosition - position)

		if distance < bestDistance {
			bestDistance = distance
			bestIdx = idx
		}
	}

	next := requests[bestIdx]

	requests = append(requests[:bestIdx], requests[bestIdx+1:]...) // pega a lista e desempacota porque espera append([]type, type, type, ...) = append([]type, ...type)

	return next, requests, nil
}
