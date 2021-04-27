package main

import (
	"fmt"

	"github.com/A9u/altilia/internal/device"
)

func main() {
	fmt.Println("----------Welcome to Altilia----------")
	fmt.Println()
	fmt.Println("Checking Battery status")
	fmt.Println("")
	device.CheckPower()
	go device.AmICharged()

	waitChannel := make(chan string)

	<-waitChannel
}
