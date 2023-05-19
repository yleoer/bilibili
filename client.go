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
	client :=  resty.New().SetTimeout(defaultTimeout)
	client.SetHeader("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0)")
	return client
}

func NewClientWithCookie(cookies []*http.Cookie) *resty.Client {
	return NewClient().SetCookies(cookies)
}

func NewRequest() *resty.Request {
	return NewClient().R()
}
