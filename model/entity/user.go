package entity

type User struct {
	UserId     uint   `json:"user_id" form:"user_id"`
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Email      string `json:"email" form:"email"`
	School     string `json:"school" form:"school"`
	UserAuthId uint   `json:"user_auth_id" form:"user_auth_id" binding:"required"`
	Roles      []*Role
	Accepted   int64
	Submitted  int64
}
