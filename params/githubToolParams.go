package params

// ToolParams 工具参数的结构体
type ToolParams struct {
	Name     string `json:"name"`
	Describe string `json:"description"`
	PushedAt string `json:"pushed_at"`
}

type ToolTagParams []struct {
	TagName string `json:"tag_name"`
}
