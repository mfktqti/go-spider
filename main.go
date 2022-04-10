package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// fetch := httpdemo("https://gorm.io/zh_CN/docs")
	// parseLink(fetch)
	// time.Sleep(3 * time.Second)

	//goquery
	//goQuery()

	gocollyDemo()
}

func parseLink(html string) {
	//替换掉空格
	html = strings.Replace(html, "\n", "", -1)
	//边栏内容块正则
	re_sidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
	sidebar := re_sidebar.FindString(html)
	re_link := regexp.MustCompile(`href="(.*?)"`)
	links := re_link.FindAllString(sidebar, -1)
	base_url := "https://gorm.io/zh_CN/docs/"
	for _, v := range links {
		s := v[6 : len(v)-1] //href="prometheus.html"
		url := base_url + s
		fmt.Printf("url: %v\n", url)
		body := httpdemo(url)
		go parseArticle(body)
	}

}

func parseArticle(body string) {
	body = strings.Replace(body, "\n", "", -1)
	re_content := regexp.MustCompile(`<div class="article">(.*?)</div>`)
	content := re_content.FindString(body)
	//fmt.Printf("content: %v\n", content)
	re_title := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)
	title := re_title.FindString(content)

	fmt.Printf("title: %v\n", title)
	title = title[42 : len(title)-5]
	fmt.Printf("title: %v\n", title)
	save(title, body)
}

func save(title, content string) {
	err := os.WriteFile("./files/"+title+".html", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func httpdemo(url string) string {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36")
	req.Header.Add("Cookie", "__gads=ID=de0a65bdae4f0cad:T=1641911930:S=ALNI_MYk8EZhj7ijqrtKG9q-sCnU80lP8w; sc_is_visitor_unique=rx12123033.1644326747.3A4FB49D8A7B4FC6020B68AA23A6A5CA.1.1.1.1.1.1.1.1.1; __utma=226521935.2058738303.1641911930.1646059603.1646059603.1; __utmz=226521935.1646059603.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); _ga_3Q0DVSGN10=GS1.1.1646837075.3.1.1646837219.43; _ga=GA1.2.2058738303.1641911930; .Cnblogs.AspNetCore.Cookies=CfDJ8AuMt_3FvyxIgNOR82PHE4lih4smmeexzC-nAuNfU7l5eo8xj1pGns89K57AipbL3DMKMlzVKBbX4RukeHk1qDcQcDvYt-qBZx3SrYRYiLJXEohy9GgoIq8urbCR9pJ4Vf9GW_3lS413GHR79ZTR_9346fLk0h55RY8qKaz6koVNmhtbcFk6GWUQ0Q0-3MP_vyXOBGrHG-fijsIz5fLWd-_RfMwuUOU3iNWnsBhnm5eUDvP6nYcHtrwwFIaqm1Gu-R9p-rSTlK3j6klM4G5V0hF-q9azj0uMkR4fza8hf803jxHxJjvZqcdD1x9voliJHU4vSklgTTopPVgkYpK9WhB4MZxcwNEeYQTpFLSZrxwbwCc4wN5JVOzv63u3At3N615Xsnv_j9xcUg3GH_YbXUHkxK_GLSa3nbodUFWCaceulsWjlZhtcNMeKWPB_9NfiAhFfHP_jcVadzUu86WoR2JXhBy794jAsuBQztg9zs3-IbtQHs2vI2ZtpBwjrXi6VmjzlpYprJKDSzvJBsKgfh-xagBg67dVT2z3PfAAJfmbiwI68OnDBH3ve0adTHSpIQ; _gid=GA1.2.1728700914.1649591356; __gpi=UID=000004b43c9ff404:T=1649591356:RT=1649591356:S=ALNI_Mbe-Lk4a224LJ6ltI1jMGndwztSQw; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1648827503,1648907088,1649130292,1649591560; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1649591560; _gat_gtag_UA_476124_1=1")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get err", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("http status code", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return ""
	}
	return string(body)
}
