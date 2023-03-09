package tool

func SendMsg(msg string) {
	appcfg := GetAppConfig()
	if appcfg.DingDing.Enable {
		msg = "Github监控:\n" + msg
		DingDingNotice(msg)
	}
}
