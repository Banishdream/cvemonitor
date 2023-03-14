package tool

import (
	"cve-monitor/define"
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"time"
)

/*
ParseBody
@describe 获取请求响应的body, 解析成结构体
@param url string "访问的URL地址"
@param method string "request请求的方法"
@param params interface{} "需要将response.body 解析成的结构体"
@return error "返回执行过程的错误信息"
*/
func ParseBody(url, method string, params interface{}) error {
	//增加错误恢复处理
	defer func() {
		if err := recover(); err != nil { // 此处进行恢复
			fmt.Printf("发生了错误,  类型: %T, err: %v\n", err, err)
		}
	}()

	appCfg := GetAppConfig()

	client := &http.Client{
		Timeout: time.Duration(define.HttpTimeout) * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	req.Header.Add("Authorization", "Bearer "+appCfg.GithubToken)
	req.Header.Set("Content-Type", "text/html; application/json; charset=utf-8")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

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

	if resp.StatusCode == http.StatusForbidden {
		fmt.Printf("url: %v, resp: %v\n", url, resp)
		fmt.Println("X-RateLimit-Reset,", resp.Header.Get("X-RateLimit-Reset"))
		fmt.Println("X-RateLimit-Used,", resp.Header.Get("X-RateLimit-Used"))
		fmt.Println("X-RateLimit-Limit,", resp.Header.Get("X-RateLimit-Limit"))
		fmt.Println("被限速了... 先休息会")
		time.Sleep(define.FuncExecInterval)
	}

	// 解析response.body
	if err = Decode(resp.Body, params); err != nil {
		log.Error("decode body err")
		return err
	}
	return err
}
