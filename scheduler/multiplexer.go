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

	for request := range requests {
		for _, requestsChannel := range requestsChannels {
			go func() {
				requestsChannel <- request
			}()
		}
	}
}
