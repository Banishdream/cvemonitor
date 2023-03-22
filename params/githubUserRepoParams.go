package params

// UserRepoParams 用户仓库的结构体
type UserRepoParams []struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Fork        bool   `json:"fork"`
	CreatedAt   string `json:"created_at"`
}
