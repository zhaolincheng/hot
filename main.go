package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hot/common/util"
	"strings"
	"time"
)

func main() {
	go getWeiBoTop()
	time.Sleep(10000000000000)
}


func getZhiHuTop() {
	fmt.Println("zhihu")
}

func getBaiDuTop()  {
	fmt.Println("baidu")
}

func getWeiBoTop() {
	url := "https://s.weibo.com/top/summary"
	response := util.HttpGet(url, nil, nil, nil)
	//str,_ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(str))
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		util.Error.Fatalln(err)
	}
	var allData []map[string]interface{}
	document.Find(".list_a li").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a span").Text()
		textLock := selection.Find("a em").Text()
		text = strings.Replace(text, textLock, "", -1)
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://s.weibo.com" + url})
		}
	})
	for index, value := range allData {
		fmt.Println(index)
		fmt.Println(value)
	}
}