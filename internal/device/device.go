package device

import (
	"fmt"
	"time"
)

func AmICharged() {
	fmt.Println("Am I charged???")
	fmt.Println()
	powerStatus := GetPower()

	percent, status := PowerStats(powerStatus)

	var duration time.Duration

	switch status {
	case FullyCharged:
		if is100Percent(percent) {
			CompletedNotifyAndRerun()
			duration = 8 * time.Hour
		} else {
			duration = 30 * time.Second
		}

	case NearlyCharged:
		duration = 30 * time.Second

	case Charging:
		if isGte95(percent) {
			duration = 2 * time.Minute
		}

	default:
		duration = 10 * time.Minute
	}

	go SleepAndRerun(duration)
	return
}

func CompletedNotifyAndRerun() {
	Notify("Charging is complete! You may now unplug the charger.")
	time.Sleep(60 * time.Second)
	Notify("Hope you have removed the charger!")
}

func SleepAndRerun(d time.Duration) {
	fmt.Printf("Going to Sleep for %v seconds\n", d.Seconds())
	time.Sleep(d)
	fmt.Println("Woke up")
	go AmICharged()
}
