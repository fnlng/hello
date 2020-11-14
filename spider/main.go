package main


import (

	"fmt"

	"io/ioutil"

	"log"

	"net/http"

	"os"

	"regexp"

	"strconv"

	"time"

)


//定义新的数据类型

type WebSite struct {

	url    string

	header map[string]string

}


//定义 WebSite get的方法

func (keyword WebSite) get_html_header() string {

	//新建客户端

	client := &http.Client{}

	//建立一个request

	req, err := http.NewRequest("GET", keyword.url, nil)

	if err != nil {

	}

	//给request添加hander

	for key, value := range keyword.header {

		req.Header.Add(key, value)

	}

	//执行request

	resp, err := client.Do(req)

	if err != nil {

		log.Println(err)

	}

	defer resp.Body.Close()

	//获取body

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		log.Println(err)

	}


	return string(body)


}

//解析parse

func parse() {

	header := map[string]string{

		"Host":                      "blog.lenconda.top",

		"Connection":                "keep-alive",

		"Cache-Control":             "max-age=0",

		"Upgrade-Insecure-Requests": "1",

		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",

		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",

		"Referer":                   "https://blog.lenconda.top",

	}



	//创建文件
	//os.MkdirAll("C://blog",0777)
	//
	//fbody,err := os.Open("C://blog/blog_spidered.html")
	//if err!=nil {
	//	log.Fatalln(err)
	//}
	//defer fbody.Close()
	//
	////博客标题
	//ftitle,err := os.Create("C://blog/blog_title.txt")
	//if err!=nil {
	//	log.Fatalln(err)
	//}
	//defer ftitle.Close()
	//ftitle.WriteString("标题\r\n")
	//
	//ftime,err := os.Create("C://blog/blog_time.txt")
	//if err!=nil {
	//	log.Fatalln(err)
	//}
	//defer ftime.Close()
	//ftime.WriteString("时间\r\n")
	//
	//ftag,err := os.Create("C://blog/blog_tag.txt")
	//if err!=nil {
	//	log.Fatalln(err)
	//}
	//defer ftag.Close()
	//ftag.WriteString("标签\r\n")
	//
	//farticle,err := os.Create("C://blog/blog_article.txt")
	//if err!=nil {
	//	log.Fatalln(err)
	//}
	//defer farticle.Close()
	//farticle.WriteString("文章\r\n")

	f,err := os.Create("C://blog/blog.html")
	if err!=nil{
		log.Fatalln(err)
	}
	defer f.Close()
	f.WriteString("<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <title>Document</title>\n</head>\n<body>")

	//循环每页解析并把结果写入excel

	for i := 1; i <= 7; i++ {

		fmt.Println("正在抓取第" + strconv.Itoa(i) + "页......")

		url := "https://blog.lenconda.top/page/" + strconv.Itoa(i) + "/"

		WebSite := &WebSite{url, header}

		html := WebSite.get_html_header()


		//标题

		pattern2 := `>(.*?)</a></h2>`

		rp2 := regexp.MustCompile(pattern2)

		find_txt2 := rp2.FindAllStringSubmatch(html, -1)

		fmt.Println(len(find_txt2))
		//时间

		pattern3 := `(.*)</time>`

		rp3 := regexp.MustCompile(pattern3)

		find_txt3 := rp3.FindAllStringSubmatch(html, -1)

		fmt.Println(len(find_txt3))
		//标签

		pattern4 := `标签(.*)`

		rp4 := regexp.MustCompile(pattern4)

		find_txt4 := rp4.FindAllStringSubmatch(html, -1)

		fmt.Println(len(find_txt4))
		//文章

		pattern5 := `(.*)</p>`

		rp5 := regexp.MustCompile(pattern5)

		find_txt5 := rp5.FindAllStringSubmatch(html,-1)

		fmt.Println(len(find_txt5))
		//  打印全部数据和写入txt文件

		for i := 0; i < len(find_txt2); i++ {

			fmt.Printf("%s\n %s %s\n %s\n", find_txt2[i][1], find_txt3[i][1], find_txt4[i][1],find_txt5[i][1])

			f.WriteString(find_txt2[i][1] + "\r\n" + find_txt3[i][1] + "\t" + find_txt4[i][1] + "\t" + "\r\n" + find_txt5[i][1] + "\r\n")

		}

	}
	f.WriteString("</body>\n</html>")

}


func main() {

	t1 := time.Now() // get current time

	parse()

	elapsed := time.Since(t1)

	fmt.Println("爬虫结束,总共耗时: ", elapsed)

}
