package tool

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
DingDingNotice
@describe 钉钉通知调用函数
@param msg string "传入要通知的消息"
*/
func DingDingNotice(msg string) {
	dtCfg := GetAppConfig().DingDing

	//请求地址模板
	mobiles := ""
	for i := 0; i < len(dtCfg.Mobiles); i++ {
		if i == 0 {
			mobiles += dtCfg.Mobiles[i]
		} else {
			mobiles += ",\n" + dtCfg.Mobiles[i]
		}
	}

	content := `{"msgtype": "text",
	"text": {"content": "` + msg + `"},
	    "at": {
	       "atMobiles": [
				` + mobiles + `
	       ],
	       "isAtAll": false
	    }
	  }`

	//创建一个请求
	req, err := http.NewRequest("POST", dtCfg.Webhook, strings.NewReader(content))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("DingTalk 发送消息成功. %v\n", string(body))
}
