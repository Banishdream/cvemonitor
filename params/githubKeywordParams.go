package params

type KeywordParams struct {
	Items []ItemConf `json:"items"`
}

type ItemConf struct {
	Name      string `json:"name"`
	HtmlUrl   string `json:"html_url"`
	CreatedAt string `json:"created_at"`
}
