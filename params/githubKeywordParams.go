package params

// KeywordParams 关键字的结构体
type KeywordParams struct {
	Items []ItemConf `json:"items"`
}

type ItemConf struct {
	Name      string `json:"name"`
	HtmlUrl   string `json:"html_url"`
	Describe  string `json:"description"`
	CreatedAt string `json:"created_at"`
}
