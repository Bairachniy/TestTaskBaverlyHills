package repo

import (
	"baverly/service/domain"
	"context"
	"database/sql"
)

type SignUpRepo struct {
	db *sql.DB
}

func NewSignUpRepo(db *sql.DB) *SignUpRepo {
	return &SignUpRepo{db: db}
}

func (s *SignUpRepo) SignUp(ctx context.Context, data domain.SignUpInfo) error {
	dao := SignUpInfo(data)

	_, err := s.db.ExecContext(ctx, `insert into customers (username, password, email) values ($1, $2, $3)`,
		dao.Username, dao.Password, dao.Email)

	return err
}
