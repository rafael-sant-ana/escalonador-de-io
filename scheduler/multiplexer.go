package scheduler

type MultiplexerHandler struct {
	schedulers []*Scheduler
}

func NewMultiplexerHandler(schedulers []*Scheduler) *MultiplexerHandler {
	return &MultiplexerHandler{
		schedulers: schedulers,
	}
}


func (h *MultiplexerHandler) ListenForAccesses(requests chan int) {
	amountListeners := len(h.schedulers)

	requestsChannels := make([]chan int, amountListeners)

	for i, scheduler := range h.schedulers {
		requestsChannels[i] = make(chan int)
		go scheduler.ListenForAccesses(requestsChannels[i])
	}

	for request := range requests {
		for _, requestsChannel := range requestsChannels {
			go func(ch chan int, req int) {
				ch <- req
			}(requestsChannel, request) // tem que passar porque essas variaveis podem ser passadas por referencia
			// e ai a iteraçao antiga poderia tentar usar o channel da nova
		}
	}
}

func (h *MultiplexerHandler) ActivateHandlers() {
	for _, scheduler := range h.schedulers {
		go scheduler.Handle()
	}
}
