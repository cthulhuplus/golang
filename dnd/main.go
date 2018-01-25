package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Tests generating multiple attributes
/*
func test() []int {
	testSlice := make([]int, 6)
	for a := 0; a < 6; a++ {
		testSlice[a] = a
//		return testSlice
	}
	return testSlice
}
*/
func main() {
	fourSide := random(1, 4)
	sixSide := random(1, 6)
	eightSide := random(1, 8)
	tenSide := random(1, 10)
	twelveSide := random(1, 12)
	twentySide := random(1, 20)
        percSide := random(1, 100)


	attSlice := make([]int, 4)
	attsum := 0

	for i := 0; i < 4; i++ {
		attSlice[i] = rand.Intn(6) + 1
	}

	  sort.Ints(attSlice)

        for _, attSlice := range attSlice[1:4] {
                attsum += attSlice
	}

	fmt.Println("D4  Roll:", fourSide)
	fmt.Println("D6  Roll:", sixSide)
	fmt.Println("D8  Roll:", eightSide)
	fmt.Println("D10 Roll:", tenSide)
	fmt.Println("D12 Roll:", twelveSide)
	fmt.Println("D20 Roll:", twentySide)
	fmt.Println("Percent Roll:", percSide, "%")

	fmt.Println("Attribute1 Rolls:", attSlice)
	fmt.Println("Attribute1 Kept:", attSlice[1:4])
	fmt.Println("Attribute1 Score:", attsum)
//	fmt.Println("Test Slice:", testSlice[1:6])
}
