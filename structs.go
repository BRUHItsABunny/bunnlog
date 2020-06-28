package bunnlog

import "log"

type BunnyLog struct {
	Logger       *log.Logger
	Verbosity    int
	StackOnErr   bool
	StackOnFatal bool
}
