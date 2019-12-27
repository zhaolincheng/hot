package top

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hot/utils"
	"strings"
)

func GetZhihuTop() {
	fmt.Println("zhihu")
}
func GetBaiduTop() {
	fmt.Println("baidu")
}

func GetWeiboTop() {
	url := "https://s.weibo.com/top/summary"
	response := utils.HttpGet(url, nil, nil, nil)
	//str,_ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(str))
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		utils.Error.Println(err)
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
