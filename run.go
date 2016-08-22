package main

import (
	"time"
	"math"
)

var numbers = map[int]Setting{
	0: {topMiddle, topRight, topLeft, bottomMiddle, bottomLeft, bottomRight},
	1: {topRight, bottomRight},
	2: {topMiddle, topRight, middle, bottomLeft, bottomMiddle},
	3: {topMiddle, topRight, middle, bottomRight, bottomMiddle},
	4: {topLeft, middle, topRight, bottomRight},
	5: {topMiddle, topLeft, middle, bottomRight, bottomMiddle},
	6: {topMiddle, topLeft, middle, bottomMiddle, bottomRight, bottomLeft},
	7: {topMiddle, topRight, bottomRight},
	8: {topMiddle, topRight, topLeft, bottomMiddle, bottomLeft, bottomRight, middle},
	9: {topMiddle, topRight, topLeft, bottomMiddle, bottomRight, middle},
}

func numberToSetting(number int) map[Digit]Setting {
	return map[Digit]Setting{
		first: numbers[digit(number, 0)],
		second: numbers[digit(number, 1)],
		third: numbers[digit(number, 2)],
		forth: numbers[digit(number, 3)],
	}
}

func digit(number int, digit int) int {
	return number / int(math.Pow(float64(10), float64(digit))) % 10
}

func main() {
	newSettings, quit := Show()

	timer := time.After(time.Duration(5) * time.Second)

	for {
		select {
		case <-timer:
			quit <- true
			<-quit
			return
		default:
			time.Now().Minute()
			newSettings <- numberToSetting(time.Now().Minute() * 100 + time.Now().Second())
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
}
