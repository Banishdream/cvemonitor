package models

type CveMonitor struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	CveName  string `xorm:"varchar(255) notnull unique" json:"cve_name"`
	PushedAt string `xorm:"varchar(255)" json:"pushed_at""`
	Url      string `xorm:"varchar(255) notnull unique" json:"url"`
}

type KeywordMonitor struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	KeywordName string `xorm:"varchar(255) notnull unique" json:"keyword_name"`
	PushedAt    string `xorm:"varchar(255)" json:"pushed_at""`
	Url         string `xorm:"varchar(255)" json:"url"`
}

type RedTeamToolsMonitor struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	ToolName string `xorm:"varchar(255) notnull unique" json:"tool_name""`
	PushedAt string `xorm:"varchar(255)" json:"pushed_at""`
	TagName  string `xorm:"varchar(255)" json:"tag_name"`
}

type UserMonitor struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	RepoName string `xorm:"varchar(255) notnull unique 'repo_name' comment('仓库名')"`
}
