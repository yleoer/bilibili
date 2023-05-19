package bilibili

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
)

const (
	FollowingsUrl = "https://api.bilibili.com/x/relation/followings"
)

func GetFollowings(client *resty.Client, uid int) {
	resp, err := client.R().SetQueryParam("mid", strconv.Itoa(uid)).Get(FollowingsUrl)
	if err != nil {
		log.Fatalf("Get followings failed: %v", err)
	}

	fmt.Println(resp)
}