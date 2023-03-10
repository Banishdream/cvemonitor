package params

// ToolParams 工具参数的结构体
type ToolParams struct {
	Name     string `json:"name"`
	PushedAt string `json:"pushed_at"`
}

type ToolTagParams []struct {
	TagName string `json:"tag_name"`
}
