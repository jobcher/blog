package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	// 写入 Markdown 文件头部
	_, err = file.WriteString("---\ntitle: " + today + " 打工人日报\ndate: " + today + "\ndraft: true\nauthor: 'jobcher'\nfeaturedImage: '/images/github.png'\nfeaturedImagePreview: '/images/github.png'\ntags: ['github']\ncategories: ['github']\nseries: ['github']\n---\n\n生活如潮水般涌来，我们都是打工人，拼搏在这个熙攘的都市。每天早出晚归，为了生计奔波不息。在这喧嚣的世界里，我们常常会忘记一些最基本的事情，比如保重身体，快乐摸鱼，以及提升自己。\n\n身体是革命的本钱。工作虽然重要，但如果我们身体垮了，一切都将付之东流。所以，无论多忙碌，我们都不能忽视保持身体的健康。多注意饮食，合理安排休息时间，保持适度的锻炼，这些都是我们保持身体健康的基本要求。身体健康，才能有充沛的精力投入工作，追求更美好的生活。\n\n快乐摸鱼是放松心情的好方法。打工人的工作压力常常让我们喘不过气来，但是，有时候放松一下也是必要的。不妨在闲暇时刻，去找一家喜欢的咖啡馆坐坐，或者约上几个好友一起逛逛街，欣赏风景。也可以在家里看一场喜剧电影，或者读一本有趣的小说。快乐的时光让我们重新充电，调整好状态，更好地迎接未来的挑战。\n\n")
	if err != nil {
		log.Fatal(err)
	}

	// 获取github热门
	get_github(md_name)
	// 获取v2ex热门
	get_v2ex(md_name)

	fmt.Println("成功生成文件")
}

// 传入 md_name 参数
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

	// 查找所有的 trending repository
	doc.Find(".cell.item").Each(func(i int, s *goquery.Selection) {
		// 提取标题和作者,title 去除span标签
		title := strings.TrimSpace(s.Find("span.item_title a").Text())
		url := strings.TrimSpace(s.Find("span.item_title a").AttrOr("href", ""))

		// 将信息以 Markdown 格式写入文件
		content := fmt.Sprintf("### %d:", i+1)
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
