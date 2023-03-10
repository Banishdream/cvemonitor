package define

import "time"

// 配置文件路径
const (
	FilePath     = "./config"
	AppFileName  = "app.json"
	ToolFileName = "tool.yaml"
)

const (
	HttpTimeout      = 15               // http request 超时时间
	HttpRetry        = 3                // http request 重试时间
	HttpMethodGET    = "GET"            // Http request 的GET方法
	HttpMethodPOST   = "POST"           // Http request 的POST方法
	RequestSleepTime = time.Second * 5  // http request 间隔时间
	FuncExecInterval = time.Second * 60 // 函数执行间隔时间
)

// WriteAll
// true:写入所有历史数据, false: 写入今天的数据, 一般第一次执行为true
const WriteAll = true

// Debug 控制数据库打印开关
const Debug = false
