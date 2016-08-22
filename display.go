package main

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

type segment byte
type Digit byte
type Setting []segment

const (
	topRight segment = iota
	bottomRight
	topLeft
	bottomLeft
	topMiddle
	middle
	bottomMiddle
	point
)

const (
	first Digit = iota
	second
	third
	forth
)

var digitPins = map[Digit]rpio.Pin{
	first:  rpio.Pin(20),
	second: rpio.Pin(16),
	third:  rpio.Pin(12),
	forth:  rpio.Pin(21),
}

var segmentPins = map[segment]rpio.Pin{
	topRight:     rpio.Pin(26),
	bottomRight:  rpio.Pin(5),
	topLeft:      rpio.Pin(19),
	bottomLeft:   rpio.Pin(22),
	topMiddle:    rpio.Pin(13),
	middle:       rpio.Pin(6),
	bottomMiddle: rpio.Pin(17),
	point:        rpio.Pin(27),
}

var AllOut = map[Digit]Setting{}

func open() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
	for _, pin := range digitPins {
		pin.Output()
	}
	for _, pin := range segmentPins {
		pin.Output()
	}
	turnDigitsOff()
}

func close() {
	turnDigitsOff()
	rpio.Close()
}

func apply(digit Digit, setting Setting) {
	digitPins[digit].Low()

	for _, seg := range setting {
		segmentPins[seg].Low()
	}
}

func turnDigitsOff() {
	for _, pin := range digitPins {
		pin.High()
	}
	for _, pin := range segmentPins {
		pin.High()
	}
}

func Show() (chan map[Digit]Setting, chan bool) {
	open()
	settings := AllOut
	newSettings := make(chan map[Digit]Setting)
	quit := make(chan bool)

	go func() {
		segmentCounter := 0
		for {
			select {
			case settings = <-newSettings:
			case <-quit:
				close()
				quit <- true
				return
			default:
				segmentCounter++
				if segmentCounter > 3 {
					segmentCounter = 0
				}
				turnDigitsOff()
				apply(Digit(segmentCounter), settings[Digit(segmentCounter)])
				time.Sleep(time.Duration(1) * time.Millisecond)
			}
		}
	}()
	return newSettings, quit
}
