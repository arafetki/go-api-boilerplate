package service

import (
	"errors"
	"time"

	"github.com/arafetki/go-echo-boilerplate/internal/db/sqlc"
	"github.com/arafetki/go-echo-boilerplate/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type userService struct {
	q *sqlc.Queries
}

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrDuplicateEmail = errors.New("email already exists")
)

func (s *userService) Create(params sqlc.CreateUserParams) error {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	err := s.q.CreateUser(ctx, params)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "users_email_key":
				return ErrDuplicateEmail
			default:
				return pgErr
			}
		}
		return err
	}
	return nil
}

func (s *userService) GetOne(id int32) (*sqlc.User, error) {
	ctx, cancel := utils.ContextWithTimeout(5 * time.Second)
	defer cancel()

	user, err := s.q.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
