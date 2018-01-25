package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Billy's Dragon Quest!")
	time.Sleep(2 * time.Second)
	fmt.Println("Created by Billy and Bill Brooks")
	time.Sleep(2 * time.Second)
	fmt.Println("Prepage for adventure!")
	time.Sleep(2 * time.Second)
	fmt.Println("You awaken in a big dark cave...")
	time.Sleep(2 * time.Second)
	fmt.Println("Do you want to go [up] towards the light or [down] towards darness?")
	cave, _ := reader.ReadString('\n')
	if strings.TrimRight(cave, "\n") == "up" {
		time.Sleep(2 * time.Second)
		fmt.Println("There is a rock in your path [push] it to escape or go back [down]?")
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("Ooops!")
	}
}
