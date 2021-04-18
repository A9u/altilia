package device

// +build darwin

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

const POWER_REGEX = `[\d]{1,2}%`

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

func Notify() {
	regex := regexp.MustCompile(POWER_REGEX)

	powerBytes := regex.Find([]byte(GetPower()))

	power := fmt.Sprintf("%s is remaining", powerBytes)

	powerCmd := fmt.Sprintf(`display notification %q with title "Power"`, power)
	cmd := exec.Command("osascript", "-e", powerCmd)

	var cmdError bytes.Buffer
	cmd.Stderr = &cmdError

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error received::", cmdError.String())
		log.Fatal(err)
	}
}
