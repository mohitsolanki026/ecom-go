package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	GetUserById(id int) (*User, error)
}

type ProductStore interface {
	GetProducts()([]Product, error)
}

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string 	`json:"image"`
	Price float64 `json:"price"`
	Quantity int 	`json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
