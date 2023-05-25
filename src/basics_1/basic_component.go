package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("number is", rand.Intn(10))

	rand.NewSource(time.Now().UnixMicro())
	randomNum := rand.Intn(101)
	fmt.Printf("today... 「 %d 」 is lucky number.\n", randomNum)

	fmt.Printf("rounded math is %g.\n", math.Round(math.Pi))

	fmt.Printf("add method test... %d.\nadd Func Method is… %d\n", add(2, 5), addFunc2(3, 4))

	swapA, swapB := swap("Hello", "world")
	fmt.Println(swapA, swapB)

	fmt.Println(noNamedReturn(17))

	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)

	n, m := 1, 2
	var cSharp, django, js = true, false, "no!"
	if cSharp {
		js = "yes!"
	}
	fmt.Println(n, m, cSharp, django, js)

	// 初期値
	var intNum int
	var floatNum float64
	var boolObj bool
	var strObj string
	fmt.Printf("%v %v %v %q\n", intNum, floatNum, boolObj, strObj)

	xNum, yNum := 3, 4
	floatNum = math.Sqrt(float64(xNum*xNum + yNum*yNum))
	zNum := uint(floatNum)
	fmt.Println(xNum, yNum, zNum)

	const Pi = 3.14
	fmt.Printf("happy %v day\n", Pi)

	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	needInt := func(x int) int {
		return x*10 + 1
	}
	needFloat := func(x float64) float64 {
		return x * 0.1 * 1
	}
	fmt.Println(needInt(Small), needFloat(Big))
}

func add(x int, y int) int {
	return x + y
}

func addFunc2(x, y int) int {
	xProductSeconds := x * 2
	yQuotantSeconds := y / 2
	return xProductSeconds + yQuotantSeconds
}

func swap(x, y string) (string, string) {
	return y, x
}

func noNamedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
