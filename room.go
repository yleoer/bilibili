package bilibili

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

const (
	InitRoomUrl = "https://api.live.bilibili.com/room/v1/Room/room_init"
)

type InitRoom struct {
	RoomId     int `json:"room_id"`
	ShortId    int `json:"short_id"`
	Uid        int `json:"uid"`
	LiveStatus int `json:"live_status"`
}

type InitRoomResult struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    InitRoom `json:"data"`
}

func GetInitRoom(client *resty.Client, pathId int) (*InitRoom, error) {
	var result InitRoomResult

	if _, err := client.R().SetResult(&result).SetQueryParam("id", strconv.Itoa(pathId)).Get(InitRoomUrl); err != nil {
		return nil, err
	}

	return &result.Data, nil
}
