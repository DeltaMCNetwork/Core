package server

import "time"

type Timer struct {
	time int64
}

func (timer *Timer) Start() {

}

func (timer *Timer) GetElapsedMilli() int64 {
	z := timer.time
	timer.Reset()
	return time.Now().UnixMilli() - z
}

func (timer *Timer) Reset() {
	timer.time = time.Now().UnixMilli()
}

func CreateTimer() *Timer {
	return &Timer{time: time.Now().UnixMilli()}
}

func getTime() int64 {
	return time.Now().UnixMilli()
}
