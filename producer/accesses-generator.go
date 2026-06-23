package producer

func ProduceRandomAccessesRequests(requests chan int) {
	go func() {
		requests <- 100
	}()
}
