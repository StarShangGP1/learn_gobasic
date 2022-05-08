package api

// User 用户信息存储文件的格式
type User struct {
	Nickname string `json:"nickname"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
