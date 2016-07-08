package cookieparser

import (
	"net/http"
	"strings"
)

// Parse raw cookie string
func Parse(raw string) []*http.Cookie {
	header := http.Header{}
	header.Add("Cookie", raw)
	request := http.Request{Header: header}
	return request.Cookies()
}

// ToString convert the cookies to a string
func ToString(cookies []*http.Cookie) string {
	var cookieStrings []string
	for _, cookie := range cookies {
		cookieStrings = append(cookieStrings, cookie.String())
	}
	return strings.Join(cookieStrings, ",")
}
