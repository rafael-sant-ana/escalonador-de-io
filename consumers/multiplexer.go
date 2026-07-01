package consumers

type MultiplexerHandler struct {
	fcfs *FCFSHandler
	sstf *SSTFHandler
}

func NewMultiplexerHandler(FCFSHandler *FCFSHandler, SSTFHandler *SSTFHandler) *MultiplexerHandler {
	return &MultiplexerHandler{
		fcfs: FCFSHandler,
		sstf: SSTFHandler,
	}
}

func (h *MultiplexerHandler) ListenForAccesses(requests chan int) {
	requestsFCFS := make(chan int)
	requestsSSTF := make(chan int)

	h.sstf.ListenForAccesses(requestsSSTF)
	h.fcfs.ListenForAccesses(requestsFCFS)
	go func() {
		for request := range requests {
			requestsFCFS <- request
			requestsSSTF <- request
		}
	}()
}
