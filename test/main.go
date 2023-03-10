package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func test(msg string) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=75083fbee21179c2e080796f3df84894ee288ae962d32e4305e010cd0b36de42`
	// content := `{"msgtype": "text",
	//         "text": {"content": "` + "告警" + msg + `"}
	//     }`

	content := `{"msgtype": "text",
    "text": {"content": "` + "告警test" + msg + `"},
         "at": {
            "atMobiles": [
              "15601614710"
            ],
            "isAtAll": false
         }
       }`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

func main() {
	test("Github监控: xxx")
}
