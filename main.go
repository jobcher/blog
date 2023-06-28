package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func translate(desc string) string {
	// 谷歌翻译API网址
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=en&tl=zh-CN&dt=t&q=" + url.QueryEscape(desc)

	fmt.Println("url:" + url)

	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// 解析JSON响应体
	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(err)
		return ""
	}

	// 提取翻译后的文本
	translation := result[0]["trans"].(string)
	fmt.Print("translation:" + translation)
	return translation
}

func main() {
	// 发起 HTTP GET 请求
	res, err := http.Get("https://www.github.com/trending")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("请求失败，状态码：%d", res.StatusCode)
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 当前日期
	today := time.Now().Format("2006-01-02")
	md_name := "github_trending_" + today + ".md"

	//判断文件具体目录是否存在
	_, err = os.Stat("content/posts/github/" + md_name)
	if err == nil {
		//文件存在
		os.Remove(md_name)
	}

	// 创建 Markdown 文件
	file, err := os.Create("content/posts/github/" + md_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入 Markdown 文件头部
	_, err = file.WriteString("---\ntitle: " + today + " GitHub 热门榜单\ndate: " + today + "\ndraft: true\nauthor: 'jobcher'\nfeaturedImage: '/images/github.png'\nfeaturedImagePreview: '/images/github.png'\ntags: ['github']\ncategories: ['github']\nseries: ['github']\n---\n\n")
	if err != nil {
		log.Fatal(err)
	}

	// 查找所有的 trending repository
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		// 提取标题和作者,title 去除span标签
		title := strings.TrimSpace(s.Find("h2.h3 a").AttrOr("href", ""))
		author := strings.TrimSpace(s.Find("span.text-normal").First().Text())
		url := strings.TrimSpace(s.Find("h2.h3 a").AttrOr("href", ""))
		desc := strings.TrimSpace(s.Find("p.col-9").Text())

		// 去除斜杠
		author = strings.Replace(author, "/", "", -1)
		fmt.Println("pre_desc:" + desc)
		desc = translate(desc)

		// 输出标题和作者
		fmt.Printf("### 排名 %d:", i+1)
		fmt.Printf("%s\n", title)
		fmt.Printf("简介: %s\n", desc)
		// fmt.Printf("URL: https://github.com%s\n", url)
		// fmt.Printf("作者: %s\n\n", author)

		// 将信息以 Markdown 格式写入文件
		content := fmt.Sprintf("### 排名 %d:", i+1)
		content += fmt.Sprintf("%s\n", title)
		content += fmt.Sprintf("- 简介: %s\n", desc)
		content += fmt.Sprintf("- URL: https://github.com%s\n", url)
		content += fmt.Sprintf("- 作者: %s\n\n", author)

		_, err = file.WriteString(content)
		if err != nil {
			log.Fatal(err)
		}

	})
	fmt.Println("成功生成文件")
}
