package define

import "time"

// 配置文件路径
const (
	FilePath     = "./config"
	AppFileName  = "app.json"
	ToolFileName = "tool.yaml"
)

// Request Params
// RequestSleepTime http请求时间间隔
// FuncExecInterval 函数循环间隔
const (
	HttpTimeout      = 15
	HttpRetry        = 3
	HttpMethodGET    = "GET"
	RequestSleepTime = time.Second * 5
	FuncExecInterval = time.Second * 60
)

// WriteAll true:写入所有历史数据, false: 写入今天的数据, 一般第一次执行为true
const WriteAll = true

// Debug 控制数据库打印开关
const Debug = false
