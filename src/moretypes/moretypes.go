package main

import (
	"fmt"
)

type StructTest1 struct {
	X int
	Y int
}

type StructTest2 []struct {
	Id   int
	Name string
}

func main() {
	pointerPrac1()
	s := StructTest1{2, 1}
	fmt.Println(StructTest1{1, 2}, s.X)
	ArrayTest()
}

func pointerPrac1() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func ArrayTest() {
	array := [6]int{1, 2, 3, 4, 5, 6}

	sum := 0
	for _, value := range array {
		sum += value
	}
	fmt.Println(sum)

	for i := 0; i < len(array); i++ {
		array[i] *= 2
	}
	fmt.Println(array)

	data := StructTest2{
		{1, "true"},
		{2, "false"},
		{3, "true"},
		{4, "true"},
		{5, "false"},
		{6, "true"},
	}

	// modelのような使い方？
	data = append(data, []struct {
		Id   int
		Name string
	}{
		{7, "name"},
		{8, "name"},
	}...)

	for _, item := range data {
		if item.Id%2 == 0 {
			fmt.Printf("Id: %d, Name: %s\n", item.Id, item.Name)
		}
	}
}
