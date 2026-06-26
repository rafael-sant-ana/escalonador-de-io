package producer

import (
	"io-scheduling/disk"
	"math/rand/v2"
	"time"
)

func ProduceRandomAccessesRequests(requests chan int, diskInfo disk.DiskInfo) {
	go func() {
		for {
			maxDiskBytes := diskInfo.GetMaxDiskBytes()
			newRequest := rand.IntN(maxDiskBytes)

			requests <- newRequest

			time.Sleep(3 * time.Second)
		}
	}()
}
