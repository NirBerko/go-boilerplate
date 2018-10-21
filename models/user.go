package models

type Identity interface {
	GetID() int
	GetName() string
	GetPhone() string
}

type User struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Phone    string `gorm:"column:phone"`
	Password string `gorm:"column:password"`
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetPhone() string {
	return u.Phone
}
