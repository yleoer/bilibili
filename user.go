package bilibili

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"strconv"
)

const (
	UserInfoUrl = "https://api.bilibili.com/x/space/wbi/acc/info"
	AccountUrl  = "https://api.bilibili.com/x/member/web/account"
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
	_, err := client.R().SetResult(&result).
		SetQueryParam("mid", strconv.Itoa(uid)).Get(UserInfoUrl)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}

	return &result.Data, nil
}

type Account struct {
	Mid      int    `json:"mid"`
	Uname    string `json:"uname"`
	Userid   string `json:"userid"`
	Sign     string `json:"sign"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	NickFree bool   `json:"nick_free"`
	Rank     string `json:"rank"`
}

type AccountResult struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    Account `json:"data"`
}

func GetAccount(client *resty.Client) (*Account, error) {
	var result AccountResult
	_, err := client.R().SetResult(&result).Get(AccountUrl)
	if err != nil {
		return nil, err
	}

	return &result.Data, nil
}
