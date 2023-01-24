package main

import (
	"github.com/getlantern/systray"
	"gopkg.in/ini.v1"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func init() {
	if _, err := os.Stat("config.ini"); os.IsNotExist(err) {
		cfg := ini.Empty()

		cfg.Section("leigod").Key("username").SetValue("19900000000")
		cfg.Section("leigod").Key("password").SetValue("password")

		_ = cfg.SaveTo("./config.ini")
	}
}

func main() {
	go systray.Run(onReady, onExit)

	cfg, err := ini.Load("config.ini")
	if err != nil {
		PushNotification("ERROR", "配置文件加载失败, 请检查同级目录下是否存在config.ini")
		os.Exit(0)
	}

	if !ProcessIsRunning("leigod.exe") {
		leigodProgramPath := cfg.Section("leigod").Key("programPath").Value()
		ex := exec.Command(leigodProgramPath)
		_ = ex.Start()
	}

	firstStartWaitingTime, err := cfg.Section("").Key("firstStartWaitingTime").Int()
	if err != nil {
		PushNotification("ERROR", "配置文件读取出错, 请检查firstStartWaitingTime是否为数字")
		os.Exit(0)
	}
	time.Sleep(time.Second * time.Duration(firstStartWaitingTime))

	fileBytes, err := os.ReadFile(cfg.Section("").Key("listPath").Value())
	if err != nil {
		PushNotification("ERROR", "游戏列表文件读取出错, 请检查同级目录下是否存在"+cfg.Section("").Key("listPath").Value())
		os.Exit(0)
	}

	gameList := strings.Split(string(fileBytes), "\n")
	username := cfg.Section("leigod").Key("username").Value()
	password := cfg.Section("leigod").Key("password").Value()

	gameExitWaitingTime, err := cfg.Section("").Key("gameExitWaitingTime").Int()
	if err != nil {
		PushNotification("ERROR", "配置文件出错, 请检查gameExitWaitingTime是否为数字")
		os.Exit(0)
	}

	for {
		if !ProcessIsRunning("leigod.exe") {
			err := LeigodPause(username, password)
			if err != nil {
				PushNotification("ERROR", err.Error())
				os.Exit(0)
			}

			PushNotification("INFO", "检测到雷神加速器已关闭，将自动暂停并退出程序")
			os.Exit(0)
		}

		if AllProcessClosed(gameList) {
			time.Sleep(time.Second * time.Duration(gameExitWaitingTime))
			if AllProcessClosed(gameList) {
				err := LeigodPause(username, password)
				if err != nil {
					PushNotification("ERROR", err.Error())
					os.Exit(0)
				}
				PushNotification("INFO", "检测到超"+strconv.Itoa(gameExitWaitingTime)+"秒未打开游戏，将自动暂停并退出程序")
				break
			}
		}
	}
}
