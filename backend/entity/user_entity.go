package entity

type User struct {
	ID       int64
	Username string
	Password string
	Salt     string
}
