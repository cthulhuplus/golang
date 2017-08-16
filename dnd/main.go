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

func main() {
	fourSide := random(1, 4)
	sixSide := random(1, 6)
	eightSide := random(1, 8)
	tenSide := random(1, 10)
	// TODO percSide TODO
	twelveSide := random(1, 12)
	twentySide := random(1, 20)

	str := make([]int, 4)
	strsum := 0

	for i := 0; i < 4; i++ {
		str[i] = rand.Intn(6) + 1
}

	sort.Ints(str)

        for _, str := range str[1:4] {
                strsum += str
}

	fmt.Println("D4  Roll:", fourSide)
	fmt.Println("D6  Roll:", sixSide)
	fmt.Println("D8  Roll:", eightSide)
	fmt.Println("D10 Roll:", tenSide)
	fmt.Println("D12 Roll:", twelveSide)
	fmt.Println("D20 Roll:", twentySide)

	fmt.Println("Strength Rolls:", str)
	fmt.Println("Strength Kept:", str[1:4])
	fmt.Println("Strength Score:", strsum)
}
