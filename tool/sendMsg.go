package tool

func SendMsg(msg string) {
	appCfg := GetAppConfig()
	if appCfg.DingDing.Enable {
		msg = "Github监控:\n" + msg
		DingDingNotice(msg)
	}

	if appCfg.EnterpriseWeChat.Enable {
		msg = "\"Github监控:\n" + msg + "\""
		EnterpriseWeChat(msg)
	}
}
