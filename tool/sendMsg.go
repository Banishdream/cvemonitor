package tool

func SendMsg(msg string) {
	appCfg := GetAppConfig()
	if appCfg.DingDing.Enable {
		newMsg := "Github监控:\n" + msg
		DingDingNotice(newMsg)
	}
	if appCfg.EnterpriseWeChat.Enable {
		newMsg := "\"Github监控:\n" + msg + "\""
		EnterpriseWeChat(newMsg)
	}
}
