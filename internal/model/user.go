package model

type User struct {
	UserId         int
	Username       string
	Password       string
	HashedPassword string
	Disabled       bool
}

func NewUser(userId int, username string, password string, hashedPassword string, disabled bool) *User {
	return &User{UserId: userId, Username: username, Password: password, HashedPassword: hashedPassword, Disabled: disabled}
}

func (u *User) PrepareOutput() *User {
	u.Password = ""
	u.HashedPassword = ""
	return u
}
