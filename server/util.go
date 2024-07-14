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

func (timer *Timer) HasTimePassed(time int64) bool {
	return timer.time+time <= int64(getTime())
}

func (timer *Timer) GetPassed() int64 {
	return int64(getTime()) - timer.time
}

func CreateTimer() *Timer {
	return &Timer{time: time.Now().UnixMilli()}
}

func getTime() int {
	return int(time.Now().UnixMilli())
}
