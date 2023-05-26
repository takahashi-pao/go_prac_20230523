package main

import (
	"fmt"
	"time"
)

func main() {
	// deferはLIFOで行われるスタック形式
	defer fmt.Println(roopSum(3))
	defer fmt.Println("defer test")
	showTime()
}

func roopSum(sum float64) float64 {
	for i := 0; i < 10; i++ {
		sum += float64(i)
		if v := 30; sum > float64(v) {
			return sum
		} else {
			v -= i
			sum += float64(v)
		}
	}
	return sum
}

func showTime() {
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("morning")
	case time.Now().Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}
}
