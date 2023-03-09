package tool

import (
	"cve-monitor/define"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
EnterpriseWeChat
@describe 企业微信通知调用函数
@param msg string "传入要通知的消息"
*/
func EnterpriseWeChat(msg string) {
	weChat := GetAppConfig().EnterpriseWeChat

	mentionUsers := ""
	for i := 0; i < len(weChat.MentionUsers); i++ {
		mentionUsers += "\"" + weChat.MentionUsers[i] + "\",\n"
	}

	mentionMobiles := ""
	for i := 0; i < len(weChat.MentionMobiles); i++ {
		mentionMobiles += "\"" + weChat.MentionMobiles[i] + "\",\n"
	}

	content := `{
    "msgtype": "text",
    "text": {
        "content":` + msg + `,
        "mentioned_list":[` + mentionUsers + `"@all"],
        "mentioned_mobile_list":[ ` + mentionMobiles + `"@all"]
    }
}`

	client := &http.Client{}
	req, err := http.NewRequest(define.HttpMethodPOST, weChat.Webhook, strings.NewReader(content))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
