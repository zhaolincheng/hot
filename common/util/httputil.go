package util

import (
	"net/http"
	"strings"
	"time"
)

var client http.Client

func init() {
	timeout := time.Duration(5 * time.Second)
	client = http.Client{
		Timeout: timeout,
	}
}

func HttpGet(url string, header, cookie, param map[string]string) *http.Response {
	paramStr := paramToString(param)
	url = url + paramStr
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		Error.Fatalln(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36")
	if len(header) != 0 {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}
	cookieStr := cookieToString(cookie)
	if cookieStr != "" {
		request.Header.Add("Cookie", cookieStr)
	}
	response, err := client.Do(request)
	if err != nil {
		Error.Fatalln(err)
	}
	return response
}

func HttpPost(url string, header, cookie map[string]string, body string) *http.Response {
	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		Error.Fatalln(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if len(header) != 0 {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}
	cookieStr := cookieToString(cookie)
	if cookieStr != "" {
		request.Header.Add("Cookie", cookieStr)
	}
	response, err := client.Do(request)
	if err != nil {
		Error.Fatalln(err)
	}
	return response
}

func paramToString(param map[string]string) string {
	if len(param) == 0 {
		return ""
	}
	paramStr := "?"
	for key, value := range param {
		paramStr = paramStr + key + "=" + value + "&"
	}
	strings.TrimSuffix(paramStr, "&")
	return paramStr

}

func cookieToString(cookie map[string]string) string {
	cookieStr := ""
	if len(cookie) == 0 {
		return cookieStr
	}
	for key, value := range cookie {
		cookieStr = cookieStr + key + "=" + value + "; "
	}
	cookieStr = strings.TrimSuffix(cookieStr, "; ")
	return cookieStr
}
