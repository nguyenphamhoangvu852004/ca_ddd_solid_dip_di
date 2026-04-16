package entities

type User struct {
	ID   string
	Name string
	Age  int
}

func (u *User) IsTeenager() bool {
	if u.Age < 16 {
		return false
	}
	return true
}

func NewUser() *User {
	return &User{}
}
