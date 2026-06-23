package producer

import (
	"math/rand/v2"
	"time"
)

func ProduceRandomAccessesRequests(requests chan int) {
	go func() {
		for {
			newRequest := rand.IntN(512)

			requests <- newRequest

			time.Sleep(3 * time.Second)
		}
	}()
}
