package producer

import (
	"fmt"
	"io-scheduling/disk"
	"math/rand/v2"
	"time"
)

func ProduceRandomAccessesRequests(requests chan int, diskInfo disk.DiskInfo) {
	go func() {
		maxDiskBytes := diskInfo.GetMaxDiskBytes()
		for range 9 { // quero criar 9 requests "estaticas"
			newRequest := rand.IntN(maxDiskBytes)
			requests <- newRequest
		}
		// porque quero que os algoritmos funcionem com 10 requests na pool, no minimo
		for {
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
