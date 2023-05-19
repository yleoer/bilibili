package bilibili

import (
	"net/http"
	"strings"
)

func ConvertCookies2Str(cookies []*http.Cookie) string {
	cookieStrs := make([]string, len(cookies))
	for i, cookie := range cookies {
		cookieStrs[i] = cookie.String()
	}
	return strings.Join(cookieStrs, ";")
}

func ConvertStr2Cookies(cookies string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", cookies)

	request := http.Request{Header: header}
	return request.Cookies()
}
