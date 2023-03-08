package params

type CveRepoParams struct {
	TotalCount        int        `json:"total_count"`
	IncompleteResults bool       `json:"incomplete_results"`
	Items             []DataItem `json:"items"`
}
type DataItem struct {
	Name      string `json:"name"`
	HtmlUrl   string `json:"html_url"`
	CreatedAt string `json:"created_at"`
}
