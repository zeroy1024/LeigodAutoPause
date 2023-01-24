package main

import (
	"golang.org/x/sys/windows"
	"strings"
)

func GetProcessList() []string {
	snapshot, err := windows.CreateToolhelp32Snapshot(0x2, 0x0)
	if err != nil {
		return nil
	}

	var processList []string
	for {
		var proc windows.ProcessEntry32
		proc.Size = 568
		err = windows.Process32Next(snapshot, &proc)
		if err != nil {
			break
		}
		processName := windows.UTF16ToString(proc.ExeFile[0:])
		if i := strings.Index(processName, ".exe"); i >= 1 {
			processList = append(processList, processName)
		}
	}

	return processList
}

func ProcessIsRunning(processName string) bool {
	runningProcessList := GetProcessList()

	for _, runningProcess := range runningProcessList {
		if runningProcess == processName {
			return true
		}
	}

	return false
}

func AllProcessClosed(gameList []string) bool {
	for _, processName := range gameList {
		if processRunning := ProcessIsRunning(processName); processRunning {
			return false
		}
	}

	return true
}
