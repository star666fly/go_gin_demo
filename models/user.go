package models

type User struct {
	Id       int
	Username string
	Password string
	Salt     string
}

func (user User) TableName() string {
	return "user"
}
