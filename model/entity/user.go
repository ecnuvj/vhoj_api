package entity

type User struct {
	UserId     uint    `json:"user_id" form:"user_id"`
	Username   string  `json:"username" form:"username"`
	Password   string  `json:"password" form:"password"`
	Email      string  `json:"email" form:"email"`
	School     string  `json:"school" form:"school"`
	UserAuthId uint    `json:"user_auth_id" form:"user_auth_id"`
	Roles      []*Role `json:"roles" form:"roles"`
	Accepted   int64   `json:"accepted" form:"accepted"`
	Submitted  int64   `json:"submitted" form:"submitted"`
	Token      string  `json:"token" form:"token"`
}

type RegisterUserInfo struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	School   string `json:"school" form:"school" binding:"required"`
}
