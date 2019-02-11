package render

import (
	"time"

	"github.com/coral/nocube/pkg/settings"
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
}

func New(s *settings.Settings) *Render {
	tt := time.Second / time.Duration(s.Global.Render.InternalTargetFPS)
	b := NewBroadcaster(1)
	return &Render{
		startTimer:          time.Now(),
		timeSinceLastUpdate: time.Now(),
		targetTickerTime:    tt,
		Update:              b,
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
