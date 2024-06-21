package server

import (
	"time"
)

type IServerLoop interface {
	Call(time int64)
}

type BasicServerLoop struct {
	IServerLoop
}

func (loop *BasicServerLoop) Call(timeBetween int64) {
	Info("Time between ticks is %dms", timeBetween)

	time.Sleep(15 * time.Millisecond)
}

func createBasicServerLoop() IServerLoop {
	return &BasicServerLoop{}
}
