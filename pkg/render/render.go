package render

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Update struct {
	FrameNumber     uint64
	TimeSinceStart  time.Duration
	TimeSinceUpdate time.Duration
}

type Render struct {
	ticker              *time.Ticker
	startTimer          time.Time
	timeSinceLastUpdate time.Time
	targetTickerTime    time.Duration
	Update              Broadcaster
	Benchmark           bool
}

func New(targetFPS int, benchmark bool) *Render {
	tt := time.Second / time.Duration(targetFPS)
	b := NewBroadcaster(1)

	return &Render{
		startTimer:          time.Now(),
		timeSinceLastUpdate: time.Now(),
		targetTickerTime:    tt,
		Update:              b,
		Benchmark:           benchmark,
	}
}

func (r *Render) Start() {
	r.ticker = time.NewTicker(r.targetTickerTime)
	log.Debug("Starting rendering loop")
	go r.onUpdate()

}

func (r *Render) onUpdate() {
	log.Debug("Updating render loop")
	u := Update{
		FrameNumber:     0,
		TimeSinceStart:  time.Since(r.startTimer),
		TimeSinceUpdate: time.Since(r.timeSinceLastUpdate),
	}

	if r.Benchmark {
		var m uint64 = 0
		ticker := time.NewTicker(5 * time.Second)
		go func() {
			for _ = range ticker.C {
				d := u.FrameNumber - m
				log.Info("System FPS: ", d/5)
				m = u.FrameNumber

			}
		}()
	}

	for {
		select {
		case <-r.ticker.C:
			u.TimeSinceStart = time.Since(r.startTimer)
			u.TimeSinceUpdate = time.Since(r.timeSinceLastUpdate)
			r.Update.Submit(u)

			u.FrameNumber++
			r.timeSinceLastUpdate = time.Now()
		}
	}
}
