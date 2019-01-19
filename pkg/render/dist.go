package render

type broadcaster struct {
	input chan Update
	reg   chan chan<- Update
	unreg chan chan<- Update

	outputs map[chan<- Update]bool
}

// The Broadcaster interface describes the main entry points to
// broadcasters.
type Broadcaster interface {
	// Register a new channel to receive broadcasts
	Register(chan<- Update)
	// Unregister a channel so that it no longer receives broadcasts.
	Unregister(chan<- Update)
	// Shut this broadcaster down.
	Close() error
	// Submit a new object to all subscribers
	Submit(Update)
}

func (b *broadcaster) broadcast(m Update) {
	for ch := range b.outputs {
		ch <- m
	}
}

func (b *broadcaster) run() {
	for {
		select {
		case m := <-b.input:
			b.broadcast(m)
		case ch, ok := <-b.reg:
			if ok {
				b.outputs[ch] = true
			} else {
				return
			}
		case ch := <-b.unreg:
			delete(b.outputs, ch)
		}
	}
}

// NewBroadcaster creates a new broadcaster with the given input
// channel buffer length.
func NewBroadcaster(buflen int) Broadcaster {
	b := &broadcaster{
		input:   make(chan Update, buflen),
		reg:     make(chan chan<- Update),
		unreg:   make(chan chan<- Update),
		outputs: make(map[chan<- Update]bool),
	}

	go b.run()

	return b
}

func (b *broadcaster) Register(newch chan<- Update) {
	b.reg <- newch
}

func (b *broadcaster) Unregister(newch chan<- Update) {
	b.unreg <- newch
}

func (b *broadcaster) Close() error {
	close(b.reg)
	return nil
}

func (b *broadcaster) Submit(m Update) {
	if b != nil {
		b.input <- m
	}
}
