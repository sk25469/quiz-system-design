package model

type UserType int

const (
	ADMIN   = 0
	STUDENT = 1
	VIEWER  = 2
)

type User struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Type  UserType `json:"user_type"`
	Score int
}

func RegisterUser(userType int, name, email string) *User {
	return &User{
		Name:  name,
		Type:  UserType(userType),
		Email: email,
		Score: 0,
	}
}
