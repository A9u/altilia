package device

// +build darwin

import (
	"bytes"
	"fmt"
	"github.com/A9u/urja"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

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
	stats := getStatInfo(GetPower())
	powerStats := getPowerStats(stats)

	Notify(strings.Join(powerStats[0:2], " "))
}

func PowerStats(statsStr string) (percent int, status string) {
	stats := getStatInfo(statsStr)
	powerStats := getPowerStats(stats)

	status = strings.TrimSpace(powerStats[1])

	percentStr := strings.TrimSpace(powerStats[0])
	percentStr = strings.Split(percentStr, "%")[0]

	percent, err := strconv.Atoi(percentStr)
	if err != nil {
		return 0, status
	}

	return percent, status
}

func Notify(message string) {
	powerCmd := fmt.Sprintf(`display notification %q with title "Power 🔋"`, message)
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

func is100Percent(percent int) bool {
	return percent == 100
}

func isGte95(percent int) bool {
	return percent >= 95
}

func isClosingBracket(r rune) bool {
	return r == ')'
}

func getStatInfo(statsStr string) (stats string) {
	return strings.FieldsFunc(statsStr, isClosingBracket)[1]
}

func getPowerStats(statsInfo string) []string {
	return strings.Split(statsInfo, ";")
}
