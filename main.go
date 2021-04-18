package main

import "fmt"
import "github.com/A9u/altilia/internal/device"

func main() {
	fmt.Println("----------Welcome to Altilia----------")
	fmt.Println()
	fmt.Println("Checking Battery status")
	fmt.Println(device.GetPower())
	fmt.Println()
	fmt.Println("Notifying")
	device.Notify()
}
