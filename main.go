package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"gopkg.in/ini.v1"
	"os"
	"strings"
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
		PushNotification("配置文件加载失败", "请检查同级目录下是否存在config.ini")
		os.Exit(0)
	}

	fileBytes, err := os.ReadFile(cfg.Section("").Key("listPath").Value())
	if err != nil {
		PushNotification("游戏列表文件读取出错", "请检查同级目录下是否存在"+cfg.Section("").Key("listPath").Value())
		os.Exit(0)
	}

	username := cfg.Section("leigod").Key("username").Value()
	password := cfg.Section("leigod").Key("password").Value()
	loginResponse, err := Login(username, password)
	if err != nil {
		PushNotification("登陆失败!", err.Error())
		os.Exit(0)
	}
	accountToken := loginResponse.Data.(map[string]interface{})["login_info"].(map[string]interface{})["account_token"].(string)

	gameList := strings.Split(string(fileBytes), "\n")

	for {
		if !processIsRunning("leigod.exe") {
			PushNotification("请先运行雷神加速器!", "没有检测到加速器")
			os.Exit(0)
		}

		if allProcessClosed(gameList) {
			pauseResponse, err := Pause(accountToken)
			if err != nil {
				PushNotification("暂停失败!", "自动暂停出错，请手动暂停: "+err.Error())
				os.Exit(0)
			}

			fmt.Println(pauseResponse.Message)

			os.Exit(0)
		}
	}

}
