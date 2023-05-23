package main

import (
	"fmt"
	"time"
)

/* メインメソッド */
func main() {
	fmt.Printf("hello, world!")
	strAndDatetime()
}

/* 文字列と現在時刻を返す */
func strAndDatetime() {
	fmt.Printf("Welcome")
	fmt.Println("to")
	fmt.Printf("the playground!")

	fmt.Println("the time is", time.Now())

	var name string = "ジョーギブソン"

	fmt.Println(name)
}
