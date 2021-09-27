package queues

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

func (t *Queue) New() {
	t.message = make(chan int)
	t.queuePass = make(chan int)
	t.queueTicket = make(chan int)
}

