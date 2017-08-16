package main

import (
	"fmt"
	"math/rand"
	"time"
	//  "."
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//func attributes() {
//	str := fourSide x 6
//	var str []int
//	str = append(str, sixSide)
//}

func main() {
	fourSide := random(1, 4)
	sixSide := random(1, 6)
	eightSide := random(1, 8)
	tenSide := random(1, 10)
	//  percSide := // TODO
	twelveSide := random(1, 12)
	twentySide := random(1, 20)

	var str []int

	str = append(str, sixSide)

	fmt.Println("D4  Roll:", fourSide)
	fmt.Println("D6  Roll:", sixSide)
	fmt.Println("D8  Roll:", eightSide)
	fmt.Println("D10 Roll:", tenSide)
	fmt.Println("D12 Roll:", twelveSide)
	fmt.Println("D20 Roll:", twentySide)

	fmt.Println("Str:", str)
}
