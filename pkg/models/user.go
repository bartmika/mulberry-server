// github.com/bartmika/mulberry-server/internal/models/user.go
package models

// User ..
type User struct {
	Name string
}

// UserRepository ..
type UserRepository interface {
	FindByID(ID int) (*User, error)
	Save(user *User) error
}
