package rapa102

import "github.com/coral/nocube/pkg/settings"

type RMan struct {
}

func New(s *settings.Settings) *RMan {
	return &RMan{}
}

func (rm *RMan) Discover() {

}
