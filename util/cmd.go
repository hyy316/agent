package util

import (
	"os/exec"
	"strings"
)

func ExecCommand(commandName string, params string) []string {
	var data []string
	cmd := exec.Command(commandName, params)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return data
	}
	strs := strings.Split(string(buf), "\n")

	for index := 0; index < len(strs)-1; index++ {
		data = append(data, strs[index])
	}
	return data
}
