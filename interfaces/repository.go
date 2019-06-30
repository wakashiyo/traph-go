package interfaces

import (
	"fmt"

	traph "github.com/wakashiyo/traph-go"
)

type Repository interface {
	Execute(query string, args ...interface{}) (Result, error)
	Query(statement string) (Rows, error)
	QueryRow(statement string, args ...interface{}) Row
}

type Rows interface {
	Scan(args ...interface{}) error
	Next() bool
	Close() error
}

type Row interface {
	Scan(args ...interface{}) error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type RepositoryImpl struct {
	repository Repository
}

func (r *RepositoryImpl) UserStore(u traph.NewUser) (*traph.User, error) {
	query := "INSERT INTO users (name, display_name, email) VALUES(?, ?, ?)"
	result, err := r.repository.Execute(query, u.Name, u.DisplayName, u.Email)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user := &traph.User{
		ID:          fmt.Sprint(id),
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Email:       u.Email,
	}
	return user, nil
}

func (r *RepositoryImpl) UserUpdate(u traph.EditedUser) (*traph.User, error) {
	query := "UPDATE users SET name = ?, display_name = ?, email = ? WHERE id = ?"
	_, err := r.repository.Execute(query, u.Name, u.DisplayName, u.Email, u.ID)
	if err != nil {
		return nil, err
	}
	user := &traph.User{
		ID:          u.ID,
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Email:       u.Email,
	}
	return user, nil
}
