package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	_, err = file.WriteString("---\ntitle: " + today + " 打工人日报\ndate: " + today + "\ndraft: true\nauthor: 'jobcher'\nfeaturedImage: '/images/wallpaper/" + today + ".jpg'\nfeaturedImagePreview: '/images/wallpaper/" + today + ".jpg'\nimages: ['/images/wallpaper/" + today + ".jpg']\ntags: ['github']\ncategories: ['github']\nseries: ['github']\n---\n\n")
	if err != nil {
		log.Fatal(err)
	}

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

	count := 0
	// 查找所有的 trending repository
	doc.Find(".Box .Box-row").Each(func(i int, s *goquery.Selection) {
		count++
		if count > 10 {
			return
		}
		// 提取标题和作者,title 去除span标签
		title := strings.TrimSpace(s.Find("h2.h3 a").AttrOr("href", ""))
		author := strings.TrimSpace(s.Find("span.text-normal").First().Text())
		url := strings.TrimSpace(s.Find("h2.h3 a").AttrOr("href", ""))
		desc := strings.TrimSpace(s.Find("p.col-9").Text())

		// 去除斜杠
		author = strings.Replace(author, "/", "", -1)
		//翻译
		queryString := desc
		result, err := translateString(queryString)
		if err != nil {
			fmt.Println("翻译失败：", err)
			return
		}
		desc = result

		// 将信息以 Markdown 格式写入文件
		content := fmt.Sprintf("### 排名 %d:", i+1)
		content += fmt.Sprintf("%s\n", title)
		content += fmt.Sprintf("- 简介: %s\n", desc)
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

// 翻译
type TranslationResponse struct {
	From         string `json:"from"`
	To           string `json:"to"`
	TransResults []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}

func translateString(queryString string) (string, error) {
	// 使用环境变量
	apiKey := os.Getenv("BAIDU_TRANSLATE_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("未设置百度翻译 API 密钥")
	}
	apiId := os.Getenv("BAIDU_TRANSLATE_API_ID")
	if apiId == "" {
		return "", fmt.Errorf("未设置百度翻译 API ID")
	}
	apiURL := "https://fanyi-api.baidu.com/api/trans/vip/translate"
	salt := "1435660288" // 随机数，这里使用固定值

	// 构建 POST 请求参数
	values := url.Values{}
	values.Set("q", queryString)
	values.Set("from", "en")
	values.Set("to", "zh")
	values.Set("appid", apiId) // 百度翻译 API 的应用ID，固定值
	sign := apiId + queryString + salt + apiKey
	fmt.Println(sign)
	values.Set("salt", salt)
	values.Set("sign", fmt.Sprintf("%x", md5.Sum([]byte(sign))))

	// 发送 POST 请求
	resp, err := http.PostForm(apiURL, values)
	if err != nil {
		return "", fmt.Errorf("请求失败：%v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败：%v", err)
	}

	// 解析 JSON 数据
	var response TranslationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("解析 JSON 失败：%v", err)
	}

	// 提取翻译结果
	if len(response.TransResults) > 0 {
		return response.TransResults[0].Dst, nil
	}

	return "", fmt.Errorf("未找到翻译结果")
}
