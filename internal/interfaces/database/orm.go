package database

type Orm interface {
	SaveUser(username string, password string) (*User, error)
	UpdateUser(username string, password string, disabled bool) (*User, error)
	FindUserByName(username string) (*User, error)
	FindUserByID(userId int) (*User, error)
}

type User struct {
	UserId         int
	Username       string
	Password       string
	HashedPassword string
	Disabled       bool
}

func NewUser(userid int, username string, password string, hashedPassword string, disabled bool) *User {
	return &User{UserId: userid, Username: username, Password: password, HashedPassword: hashedPassword, Disabled: disabled}
}
