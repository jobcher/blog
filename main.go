package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 当前日期
	today := time.Now().Format("2006-01-02")
	md_name := "github_trending_" + today + ".md"

	//判断文件是否存在
	_, err := os.Stat("content/posts/github/" + md_name)
	if err == nil {
		fmt.Println("文件已存在")
		os.Exit(0)
	}

	// 创建 Markdown 文件
	file, err := os.Create("content/posts/github/" + md_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	downloadBingWallpaper()

	// 写入 Markdown 文件头部
	_, err = file.WriteString("---\ntitle: " + today + " 打工人日报\ndate: " + today + "\ndraft: true\nauthor: 'jobcher'\nfeaturedImage: '/images/wallpaper/" + today + ".jpg'\nfeaturedImagePreview: '/images/wallpaper/" + today + ".jpg'\ntags: ['github']\ncategories: ['github']\nseries: ['github']\n---\n\n")
	if err != nil {
		log.Fatal(err)
	}
	// 每日英语
	iciba(md_name)
	// 获取微博热搜
	get_weibo(md_name)
	// 获取github热门
	get_github(md_name)
	// 获取v2ex热门
	get_v2ex(md_name)

	fmt.Println("成功生成文件")
}

func get_weibo(md_name string) {
	//写入标题
	file, err := os.OpenFile("content/posts/github/"+md_name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("## 微博热搜榜\n\n")

	// 发起 HTTP GET 请求
	res, err := http.Get("https://tophub.today/n/KqndgxeLl9")
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

	count := 0
	// 查找所有的热搜
	doc.Find(".table tbody tr").Each(func(i int, s *goquery.Selection) {
		count++
		if count > 20 {
			return
		}
		// 提取标题和url
		title := strings.TrimSpace(s.Find("td a").Text())
		url := strings.TrimSpace(s.Find("td a").Text())

		title = strings.Replace(title, "", "", -1)
		url = strings.Replace(url, "", "", -1)
		url = strings.Replace(url, " ", "", -1)

		// 将信息以 Markdown 格式写入文件
		content := fmt.Sprintf("#### 排名 %d.", i+1)
		content += fmt.Sprintf("[%s]", title)
		content += fmt.Sprintf("(https://s.weibo.com/weibo?q=%s)\n", url)

		fmt.Println(content)

		// 写入 Markdown 文件
		file, err := os.OpenFile("content/posts/github/"+md_name, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(content)
	})
}

func get_github(md_name string) {
	//写入标题
	file, err := os.OpenFile("content/posts/github/"+md_name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("## GitHub 热门榜单\n\n")

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

	// 查找所有的 trending repository
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		// 提取标题和作者,title 去除span标签
		title := strings.TrimSpace(s.Find("h2.h3 a").AttrOr("href", ""))
		author := strings.TrimSpace(s.Find("span.text-normal").First().Text())
		url := strings.TrimSpace(s.Find("h2.h3 a").AttrOr("href", ""))
		desc := strings.TrimSpace(s.Find("p.col-9").Text())

		// 去除斜杠
		author = strings.Replace(author, "/", "", -1)

		// 将信息以 Markdown 格式写入文件
		content := fmt.Sprintf("### 排名 %d:", i+1)
		content += fmt.Sprintf("%s\n", title)
		content += fmt.Sprintf("- Description: %s\n", desc)
		content += fmt.Sprintf("- URL: https://github.com%s\n", url)
		content += fmt.Sprintf("- 作者: %s\n\n", author)

		fmt.Println(content)

		// 写入 Markdown 文件
		file, err := os.OpenFile("content/posts/github/"+md_name, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(content)
	})
}

func get_v2ex(md_name string) {
	//写入标题
	file, err := os.OpenFile("content/posts/github/"+md_name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("## v2ex 热门帖子\n\n")

	// 发起 HTTP GET 请求
	res, err := http.Get("https://www.v2ex.com/?tab=hot")
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

	count := 0
	// 查找所有的 trending repository
	doc.Find(".cell.item").Each(func(i int, s *goquery.Selection) {
		count++
		if count > 20 {
			return
		}
		// 提取标题和作者,title 去除span标签
		title := strings.TrimSpace(s.Find("span.item_title a").Text())
		url := strings.TrimSpace(s.Find("span.item_title a").AttrOr("href", ""))

		title = strings.Replace(title, " ", "", -1)
		url = strings.Replace(url, " ", "", -1)

		// 将信息以 Markdown 格式写入文件
		content := fmt.Sprintf("#### %d.", i+1)
		content += fmt.Sprintf("[%s]", title)
		content += fmt.Sprintf("(https://www.v2ex.com%s)\n", url)

		fmt.Println(content)

		// 写入 Markdown 文件
		file, err := os.OpenFile("content/posts/github/"+md_name, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(content)
	})
}

type BingResponse struct {
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
}

func downloadBingWallpaper() {
	// 获取当前日期
	currentTime := time.Now()
	dateString := currentTime.Format("2006-01-02")

	// 指定保存目录
	saveDirectory := "static/images/wallpaper/"

	// 构建保存文件路径
	savePath := filepath.Join(saveDirectory, dateString+".jpg")

	// 发起 HTTP 请求获取 Bing 每日壁纸信息
	response, err := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US")
	if err != nil {
		fmt.Println("无法获取壁纸信息:", err)
		return
	}
	defer response.Body.Close()

	// 解析 JSON 数据
	var bingResponse BingResponse
	err = json.NewDecoder(response.Body).Decode(&bingResponse)
	if err != nil {
		fmt.Println("解析壁纸信息失败:", err)
		return
	}

	if len(bingResponse.Images) == 0 {
		fmt.Println("未找到壁纸信息")
		return
	}

	// 获取壁纸 URL
	imageURL := "https://www.bing.com" + bingResponse.Images[0].URL

	// 发起 HTTP 请求下载壁纸
	imageResponse, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("无法下载壁纸:", err)
		return
	}
	defer imageResponse.Body.Close()

	// 创建保存文件
	file, err := os.Create(savePath)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	// 将壁纸内容保存到文件
	_, err = io.Copy(file, imageResponse.Body)
	if err != nil {
		fmt.Println("保存壁纸失败:", err)
		return
	}

	fmt.Println("壁纸已成功保存到:", savePath)

}

func iciba(md_name string) {
	// 发起 HTTP GET 请求
	url := "http://news.iciba.com/views/dailysentence/daily.html#!/detail/title/" + time.Now().Format("2006-01-02")
	fmt.Println(url)
	res, err := http.Get(url)
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

	// 获取detail-content-en
	doc.Find(".detail-box").Each(func(i int, s *goquery.Selection) {
		detail_en := strings.TrimSpace(s.Find(".detail-content-en").Text())
		detail_zh := strings.TrimSpace(s.Find(".detail-content-zh").Text())

		fmt.Println(detail_en)
		fmt.Println(detail_zh)

		// 将信息以 Markdown 格式写入文件
		content := fmt.Sprintf(">%s\n", detail_en)
		content += fmt.Sprintf(">%s\n\n", detail_zh)

		fmt.Println(content)

		// 写入 Markdown 文件
		file, err := os.OpenFile("content/posts/github/"+md_name, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString(content)
	})
}
