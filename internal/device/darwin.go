package device

// +build darwin

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func GetPower() {
	cmd := exec.Command("pmset", "-g", "ps")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
