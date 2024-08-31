package user

import (
	"database/sql"
	"fmt"

	"github.com/mohitsolanki026/econ-go/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User,  error){
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?",email)
	if err != nil {
		return nil, err
		
	}
	u := new(types.User)
	for rows.Next(){
		u, err = scanRowIntUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowIntUser(rows *sql.Rows) (*types.User, error){
	user := new(types.User)

	err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO Users(first_name, last_name,email,password) VALUES(?,?,?,?)", user.FirstName,user.LastName, user.Email, user.Password)
	
	if err != nil {
		return err
	}

	return nil

}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows,err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil,err
	}

	u := new(types.User)
	for rows.Next(){
		u,err = scanRowIntUser(rows)
		if err != nil{
			return nil,err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	} 

	return u, nil
}