package types

import "time"

type UserStore interface {
	GetUserByEmail(email string)(*User ,error)
	CreateUser(user *User) error
	GetUserById(id int) (*User, error)
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=32"`
}

type LoginUser struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
