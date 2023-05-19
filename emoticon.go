package bilibili

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

const (
	GetEmoticonUrl = "https://api.live.bilibili.com/xlive/web-ucenter/v2/emoticon/GetEmoticons"
)

type Emoticon struct {
	BulgeDisplay      int    `json:"bulge_display"`
	Descript          string `json:"descript"`
	Emoji             string `json:"emoji"`
	EmoticonId        int    `json:"emoticon_id"`
	EmoticonUnique    string `json:"emoticon_unique"`
	EmoticonValueType int    `json:"emoticon_value_type"`
	Height            int    `json:"height"`
	Width             int    `json:"width"`
	Identity          int    `json:"identity"`
	InPlayerArea      int    `json:"in_player_area"`
	IsDynamic         int    `json:"is_dynamic"`
	Perm              int    `json:"perm"`
	UnlockNeedGift    int    `json:"unlock_need_gift"`
	UnlockNeedLevel   int    `json:"unlock_need_level"`
	UnlockShowColor   string `json:"unlock_show_color"`
	UnlockShowImage   string `json:"unlock_show_image"`
	UnlockShowText    string `json:"unlock_show_text"`
	Url               string `json:"url"`
}

type EmoticonPackage struct {
	CurrentCover string     `json:"current_cover"`
	Emoticons    []Emoticon `json:"emoticons"`
	PkgDescript  string     `json:"pkg_descript"`
	PkgId        int        `json:"pkg_id"`
	PkgName      string     `json:"pkg_name"`
	PkgPerm      int        `json:"pkg_perm"`
	PkgType      int        `json:"pkg_type"`
}

type EmoticonResult struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    struct{
		FansBrand int `json:"fans_brand"`
		Data []EmoticonPackage `json:"data"`
	} `json:"data"`
}

func GetEmoticonPackage(client *resty.Client, roomId int) ([]EmoticonPackage, error) {
	var result EmoticonResult
	_, err := client.R().SetResult(&result).SetQueryParam("platform", "pc").SetQueryParam("room_id", strconv.Itoa(roomId)).Get(GetEmoticonUrl)

	return result.Data.Data, err
}
