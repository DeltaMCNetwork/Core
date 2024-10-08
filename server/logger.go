package server

import (
	"fmt"
	"os"
	"time"
)

func getTimeFormat() string {
	return time.Now().Format("15:04:05")
}

// Debug .
func Debug(args ...any) {
	fmt.Printf("["+getTimeFormat()+"] [\x1B[94mDEBUG\x1B[0m] "+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
}

// Info .
func Info(args ...any) {
	fmt.Printf("["+getTimeFormat()+"] [\x1B[92mINFO\x1B[0m] "+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
}

// Warn .
func Warn(args ...any) {
	fmt.Printf("["+getTimeFormat()+"] [\x1B[33mWARN\x1B[0m] \x1B[33m"+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
}

// Error .
func Error(args ...any) {
	fmt.Printf("["+getTimeFormat()+"] [\x1B[91mERROR\x1B[0m] \x1B[91m"+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)
}

// Fatal .
func Fatal(args ...any) {
	fmt.Printf("["+getTimeFormat()+"] [\x1B[91m\x1B[1mFATAL\x1B[0m] \x1B[91m\x1B[1m"+fmt.Sprintf("%v\x1B[0m\n", args[0]), args[1:]...)

	os.Exit(1)
}
