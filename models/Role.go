package models

type Role struct {
	Id   int
	Name string
}

func (role Role) TableName() string {
	return "role"
}
