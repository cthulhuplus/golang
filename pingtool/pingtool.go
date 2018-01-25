package main

import (
	//  "fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "-c 1", "uehara-ke.com")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
