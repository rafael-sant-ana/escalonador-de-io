package producer

import (
	"fmt"
	"io-scheduling/disk"
	"math/rand/v2"
	"time"
)

func ProduceRandomAccessesRequests(requests chan int, diskInfo disk.DiskInfo) {
	go func() {
		for {
			maxDiskBytes := diskInfo.GetMaxDiskBytes()
			newRequest := rand.IntN(maxDiskBytes)

			producerLog(newRequest)
			requests <- newRequest

			time.Sleep(3 * time.Second)
		}
	}()
}

func producerLog(a ...any) {
	fmt.Print("[PRODUCER] ")
	fmt.Println(a...)
}
