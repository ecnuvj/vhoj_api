package entity

type User struct {
	UserId     uint
	Username   string
	Password   string
	Email      string
	School     string
	UserAuthId uint
	Roles      []*Role
	Accepted   int64
	Submitted  int64
}
