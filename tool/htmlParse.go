package tool

import (
	"cve-monitor/define"
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
	client := &http.Client{
		Timeout: time.Duration(define.HttpTimeout) * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")

	if err != nil {
		log.Fatalf("http NewRequest err->%v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("do client err->%v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read resp err->%v", err)
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
