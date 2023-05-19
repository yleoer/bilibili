package bilibili

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const (
	defaultTimeout = 20 * time.Second
)

func NewClient() *resty.Client {
	return resty.New().SetTimeout(defaultTimeout)
}

func NewClientWithCookie(cookies []*http.Cookie) *resty.Client {
	return NewClient().SetCookies(cookies)
}

func NewRequest() *resty.Request {
	return NewClient().R()
}
