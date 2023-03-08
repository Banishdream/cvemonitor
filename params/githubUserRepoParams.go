package params

type UserRepoParams []struct {
	FullName  string `json:"full_name"`
	Fork      bool   `json:"fork"`
	CreatedAt string `json:"created_at"`
}
