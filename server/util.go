package server

import (
	"fmt"
	"strings"
	"time"
)

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

func ReadVarInt(data []byte) (int32, int) {
	var num int
	var res int32

	for {
		val := data[num]
		tmp := int32(val)
		res |= (tmp & 0x7F) << uint(num*7)

		if num++; num > 5 {
			panic("Value too big!")
		}

		if tmp&0x80 != 0x80 {
			break
		}
	}

	return res, num
}

func GetVarIntBytes(value int) []byte {
	data := make([]byte, 0)

	for (value & -128) != 0 {
		data = append(data, byte(value&127|128))
		value = value >> 7
	}

	return data
}

func Stringify(value string, values ...any) string {
	var fixed = strings.ReplaceAll(value, "&", "§")

	return fmt.Sprintf(fixed, values...)
}
