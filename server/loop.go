package server

import (
	"time"

	"github.com/charmbracelet/log"
)

type IServerLoop interface {
	Call(time int64)
}

type BasicServerLoop struct {
	IServerLoop
}

func (loop *BasicServerLoop) Call(timeBetween int64) {
	log.Info("time", timeBetween)

	time.Sleep(15 * time.Millisecond)
}

func createBasicServerLoop() IServerLoop {
	return &BasicServerLoop{}
}
