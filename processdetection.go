package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func processIsRunning(processName string) bool {
	process := strings.Split(processName, ".")

	buffer := bytes.Buffer{}
	wmic := exec.Command("wmic", "process", "get", "name,executablepath")
	wmic.Stdout = &buffer
	err := wmic.Run()
	if err != nil {
		return false
	}

	findstr := exec.Command("findstr", process[0])
	findstr.Stdin = &buffer
	data, err := findstr.CombinedOutput()
	if err != nil {
		return false
	}

	if !strings.Contains(string(data), processName) {
		return false
	}

	return true
}

func allProcessClosed(gameList []string) bool {
	for _, processName := range gameList {
		if processRunning := processIsRunning(processName); processRunning {
			return false
		}
	}

	return true
}
