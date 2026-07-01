package consumers

type MultiplexerHandler struct {
	fcfs *FCFSHandler
}

func NewMultiplexerHandler(FCFSHandler *FCFSHandler) *MultiplexerHandler {
	return &MultiplexerHandler{
		fcfs: FCFSHandler,
	}
}

func (h *MultiplexerHandler) ListenForAccesses(requests chan int) {
	requestsFCFS := make(chan int)

	h.fcfs.ListenForAccesses(requestsFCFS)
	go func() {
		for request := range requests {
			requestsFCFS <- request
		}
	}()
}
