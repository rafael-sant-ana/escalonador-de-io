package main

import (
	"io-scheduling/disk"
	"io-scheduling/producer"
	"io-scheduling/scheduler"
	"io-scheduling/scheduler/strategies"
)

func main() {
	requestsChannel := make(chan int)

	initialPosition := 256
	maxDiskBytes := 512
	diskInfo := disk.NewDiskInfo(
		initialPosition,
		maxDiskBytes,
	)

	FCFSStrategy := strategies.NewFCFSStrategy()
	SSTFStrategy := strategies.NewSSTFStrategy()

	FCFSScheduler := scheduler.NewScheduler(
		diskInfo,
		initialPosition,
		FCFSStrategy,
	)

	SSTFScheduler := scheduler.NewScheduler(
		diskInfo,
		initialPosition,
		SSTFStrategy,
	)

	schedulers := []*scheduler.Scheduler{
		FCFSScheduler,
		SSTFScheduler,
	}

	multiplexer := scheduler.NewMultiplexerHandler(schedulers)

	producer.ProduceRandomAccessesRequests(requestsChannel, diskInfo)

	go multiplexer.ListenForAccesses(requestsChannel)
	multiplexer.ActivateHandlers()

	select {} // Deixa a main aberta pra smp
	// porque: o select eh usado para esperar eventos de goroutines. como nao temos casos, estamos esperando pra smp.
}
