package tool

import (
	"cve-monitor/define"
	"fmt"
	"github.com/prometheus/common/log"
	"io"
	"net/http"
	"regexp"
	"time"
)

/*
IsExistCVE
@describe 解析 html 判断title中是否包含关键字 CVE -ERROR: Couldn't find
@param url string "要获取的url地址"
*/
func IsExistCVE(url string) bool {
	//增加错误恢复处理
	defer func() {
		if err := recover(); err != nil { // 此处进行恢复
			fmt.Printf("发生了错误,  类型: %T, err: %v\n", err, err)
		}
	}()

	client := &http.Client{
		Timeout: time.Duration(define.HttpTimeout) * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	if err != nil {
		log.Error("http NewRequest err->%v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("do client err->%v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("read resp err->%v", err)
	}

	defer resp.Body.Close()

	// 解析title
	title := `<title>CVE -
ERROR: Couldn't find (.*?)
</title>`
	rp := regexp.MustCompile(title)
	txt := rp.FindAllStringSubmatch(string(body), -1)

	//fmt.Println(txt)
	if len(txt) != 0 {
		return false
	}
	return true
}
