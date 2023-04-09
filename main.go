package main

import (
	"github.com/go-vgo/robotgo"
)

var active = false

func main() {
	go addKeyListen("z")
	run()
}

func addKeyListen(key string) {
	for {
		if ok := robotgo.AddEvent(key); ok {
			active = !active
		}
	}
}

func run() {
	for {
		if !active {
			continue
		}
		robotgo.Click()
		robotgo.MicroSleep(10)
		robotgo.KeyDown("enter")
		robotgo.MicroSleep(10)
		robotgo.KeyUp("enter")
	}
}
