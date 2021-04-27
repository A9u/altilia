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

	if isCharged(status) {
		if is100Percent(percent) {
			go CompletedNotifyAndRerun()
			return
		} else {
			go SleepAndRerun(30 * time.Second)
			return
		}
	}

	if isFinishingCharge(status) {
		go SleepAndRerun(30 * time.Second)
		return
	}

	if isCharging(status) {
		if isGte95(percent) {
			go SleepAndRerun(2 * time.Minute)
			return
		}
	}

	go SleepAndRerun(10 * time.Minute)
	return
}

func CompletedNotifyAndRerun() {
	Notify("Charging is complete! You may now unplug the charger.")
	time.Sleep(60 * time.Second)
	Notify("Hope you have removed the charger!")
	SleepAndRerun(8 * time.Hour)
}

func SleepAndRerun(d time.Duration) {
	fmt.Println("Going to Sleep")
	time.Sleep(d)
	fmt.Println("Woke up")
	go AmICharged()
}
