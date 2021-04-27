package device

// +build darwin

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const POWER_REGEX = `[\d]{1,3}%;\s[\S]+[\s(\S)+]*;`

func GetPower() string {
	cmd := exec.Command("pmset", "-g", "ps")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}

func CheckPower() {
	regex := regexp.MustCompile(POWER_REGEX)

	powerBytes := regex.Find([]byte(GetPower()))

	power := fmt.Sprintf("%s", powerBytes)

	Notify(power)
}

func PowerStats(statsStr string) (string, string) {
	regex := regexp.MustCompile(POWER_REGEX)

	powerBytes := regex.Find([]byte(statsStr))

	powerStats := fmt.Sprintf("%s", powerBytes)

	stats := strings.Split(powerStats, ";")

	percent := strings.Trim(stats[0], " ")
	status := strings.Trim(stats[1], " ")

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

func isCharged(status string) bool {
	return strings.EqualFold(status, "charged")
}

func is100Percent(percent string) bool {
	return strings.EqualFold(percent, "100%")
}

func isFinishingCharge(status string) bool {
	return strings.EqualFold(status, "finishing charge")
}

func isCharging(status string) bool {
	return strings.EqualFold(status, "charging")
}

func isGte95(percent string) bool {
	pct, err := strconv.Atoi(strings.Split(percent, "%")[0])
	if err != nil {
		return pct >= 95
	}

	return false
}
