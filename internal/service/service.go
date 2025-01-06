package service

import "github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"

type Service struct {
	Users interface {
		Create(params sqlc.CreateUserParams) error
		GetOne(id int32) (*sqlc.User, error)
	}
}

func New(q *sqlc.Queries) *Service {
	return &Service{
		Users: &userService{q},
	}
}
