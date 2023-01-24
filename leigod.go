package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Response struct {
	Code    uint        `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

/*type LoginResponse struct {
	LoginInfo struct {
		AccountToken string `json:"account_token"`
		ExpiryTime   string `json:"expiry_time"`
		NNToken      string `json:"nn_token"`
	} `json:"login_info"`
	UserInfo struct {
		Nickname   string `json:"nickname"`
		Avatar     string `json:"avatar"`
		Email      string `json:"email"`
		Mobile     string `json:"mobile"`
		RegionCode uint   `json:"region_code"`
	} `json:"user_info"`
}*/

func Login(username, password string) (Response, error) {
	url := "https://webapi.leigod.com/api/auth/login"
	data := make(map[string]interface{})
	data["account_token"] = nil
	data["country_code"] = 86
	data["lang"] = "zh_CN"
	data["mobile_num"] = username
	data["os_type"] = 4
	data["password"] = passwordMD5(password)
	data["region_code"] = 1
	data["sem_ad_img_url"] = map[string]string{
		"btn_yrl": "",
		"url":     "",
	}
	data["src_channel"] = "guanwang"
	data["username"] = username

	bytesData, err := json.Marshal(data)
	if err != nil {
		return Response{}, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		return Response{}, err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.61")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return Response{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var responseBody Response
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return Response{}, err
	}

	return responseBody, nil
}

func Pause(accountToken string) (Response, error) {
	url := "https://webapi.leigod.com/api/user/pause"

	data := make(map[string]interface{})
	data["lang"] = "zh_CN"
	data["account_token"] = accountToken
	bytesData, err := json.Marshal(data)
	if err != nil {
		return Response{}, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		return Response{}, err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.61")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return Response{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var responseBody Response
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return Response{}, err
	}

	return responseBody, nil
}

func LeigodPause(username, password string) error {
	loginResponse, err := Login(username, password)
	if err != nil {
		return errors.New("登陆失败: " + err.Error())
	}

	accountToken := loginResponse.Data.(map[string]interface{})["login_info"].(map[string]interface{})["account_token"].(string)
	_, err = Pause(accountToken)
	if err != nil {

		return errors.New("暂停失败: " + err.Error())
	}

	return nil
}
