package device

// +build darwin

import (
	"bytes"
	"fmt"
	"github.com/A9u/urja"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const POWER_REGEX = `[\d]{1,3}%;\s[\S]+[\s(\S)+]*;`

var regex = regexp.MustCompile(POWER_REGEX)

const (
	FullyCharged  = "charged"
	NearlyCharged = "finishing charge"
	Charging      = "charging"
)

func GetPower() string {
	output, err := urja.GetBatteryStatus()
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func CheckPower() {
	powerBytes := regex.Find([]byte(GetPower()))

	power := fmt.Sprintf("%s", powerBytes)

	Notify(power)
}

func PowerStats(statsStr string) (percent string, status string) {
	powerBytes := regex.Find([]byte(statsStr))

	powerStats := fmt.Sprintf("%s", powerBytes)

	stats := strings.Split(powerStats, ";")

	percent = strings.Trim(stats[0], " ")
	status = strings.Trim(stats[1], " ")

	return percent, status
}

func Notify(message string) {
	powerCmd := fmt.Sprintf(`display notification %q with title "Power"`, message)
	cmd := exec.Command("osascript", "-e", powerCmd)

	var cmdError bytes.Buffer
	cmd.Stderr = &cmdError

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error received::", cmdError.String())
		log.Fatal(err)
	}
}

// utils

func is100Percent(percent string) bool {
	return strings.EqualFold(percent, "100%")
}

func isGte95(percent string) bool {
	pct, err := strconv.Atoi(strings.Split(percent, "%")[0])

	if err != nil {
		return false
	}

	return pct >= 95
}
