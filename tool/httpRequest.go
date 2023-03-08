package tool

import (
	"cve-monitor/define"
	"github.com/prometheus/common/log"
	"net/http"
	"time"
)

/*
ParseBody : 获取请求响应的body, 解析成结构体
url: 访问的地址
method: 访问的方法
params: 需要解析成的结构体
*/
func ParseBody(url, method string, params interface{}) error {
	appCfg := GetAppConfig()

	client := &http.Client{
		Timeout: time.Duration(define.HttpTimeout) * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	req.Header.Set("authorization", "Bearer "+appCfg.GithubToken)
	req.Header.Set("Content-Type", "text/html; application/json; charset=utf-8")
	//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	//req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	if err != nil {
		log.Error("http NewRequest err")
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error("do client err")
		return err
	}

	defer resp.Body.Close()
	//fmt.Printf("url: %v, resp: %v\n", url, resp)
	// 解析response.body
	if err = Decode(resp.Body, params); err != nil {
		log.Error("decode body err")
		return err
	}
	return err
}
