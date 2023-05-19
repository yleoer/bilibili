package main

import (
	"bilibili"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"os"
)

const cookiesFilePath = "cookies.txt"

func main() {
	var client *resty.Client
	var err error

	if FileExists(cookiesFilePath) {
		data, err := os.ReadFile(cookiesFilePath)
		if err != nil {
			log.Fatalf("Read cookies failed: %v", err)
		}
		cookies := bilibili.ConvertStr2Cookies(string(data))
		client = bilibili.NewClientWithCookie(cookies)
	} else {
		fmt.Println("未找到 cookies，重新生成二维码")
		client, err = bilibili.LoginByQRCode()
		if err != nil {
			log.Fatalf("LoginByQRCode failed: %v", err)
		}

		cookiesStr := bilibili.ConvertCookies2Str(client.Cookies)
		if err = os.WriteFile(cookiesFilePath, []byte(cookiesStr), 0644); err != nil {
			log.Fatalf("Store cookies failed, cookies: %s, err: %v", cookiesStr, err)
		}
	}

	initRoom, err := bilibili.GetInitRoom(client, 502)
	if err != nil {
		log.Fatalf("Get init room failed: %v", err)
	}

	userInfo, err := bilibili.GetUserInfo(client, initRoom.Uid)
	if err != nil {
		log.Fatalf("Get user info failed: %v", err)
	}

	fmt.Println(*userInfo)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}

	return !os.IsNotExist(err)
}
