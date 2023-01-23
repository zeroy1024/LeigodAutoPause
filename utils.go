package main

import (
	"crypto/md5"
	"fmt"
	"github.com/go-toast/toast"
)

func passwordMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func PushNotification(title, message string) {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   title,
		Message: message,
	}
	_ = notification.Push()
}
