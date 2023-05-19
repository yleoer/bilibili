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

	account, err := bilibili.GetAccount(client)
	if err != nil {
		log.Fatalf("Get account failed: %v", err)
	}

	bilibili.GetFollowings(client, account.Mid)
	return

	var data = make(map[string]map[string]map[string]string)

	initRoom, err := bilibili.GetInitRoom(client, 502)
	if err != nil {
		log.Fatalf("Get init room failed: %v", err)
	}

	userInfo, err := bilibili.GetUserInfo(client, initRoom.Uid)
	if err != nil {
		log.Fatalf("Get user info failed: %v", err)
	}

	data[userInfo.Name] = make(map[string]map[string]string)

	fmt.Println(initRoom.RoomId)
	emoticonPackages, err := bilibili.GetEmoticonPackage(client, initRoom.RoomId)
	if err != nil {
		log.Fatalf("Get emoticon failed: %v", err)
	}

	for _, emoticonPackage := range emoticonPackages {
		if emoticonPackage.PkgName == "通用表情" {
			continue
		}

		emoticonMap := make(map[string]string, len(emoticonPackage.Emoticons))
		for _, emoticon := range emoticonPackage.Emoticons {
			emoticonMap[emoticon.Emoji] = emoticon.Url
		}

		data[userInfo.Name][emoticonPackage.PkgName] = emoticonMap
	}

	fmt.Println(data)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}

	return !os.IsNotExist(err)
}
