package bilibili

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

const (
	UserInfoUrl = "https://api.bilibili.com/x/space/acc/info"
)

type UserInfo struct {
	Mid      int    `json:"mid"`
	Name     string `json:"name"`
	Sex      string `json:"sex"`
	Face     string `json:"face"`
	Sign     string `json:"sign"`
	Rank     int    `json:"rank"`
	Level    int    `json:"level"`
	Jointime int    `json:"jointime"`
	Moral    int    `json:"moral"`
	Silence  int    `json:"silence"`
}

type UserInfoResult struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    UserInfo `json:"data"`
}

func GetUserInfo(client *resty.Client, uid int) (*UserInfo, error) {
	if len(client.Cookies) == 0 {
		return nil, errors.New("缺少 cookies")
	}

	var result UserInfoResult
	resp, err := client.R().SetResult(&result).
		SetQueryParam("mid", strconv.Itoa(uid)).Get(UserInfoUrl)
	if err != nil {
			return nil, err
	}

	fmt.Println(resp, result)

	return &result.Data, nil
}
