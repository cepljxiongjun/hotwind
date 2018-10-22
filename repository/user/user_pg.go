package user

import (
	"context"
	"github.com/cepljxiongjun/hotwind/helper"
	"github.com/cepljxiongjun/hotwind/models"
	"github.com/cepljxiongjun/hotwind/repository"
	"github.com/jmoiron/sqlx"
)

type PgUserRepo struct {
	Conn *sqlx.DB
}

func NewPgUserRepo(Conn *sqlx.DB) repository.UserRepo {
	return &PgUserRepo{
		Conn: Conn,
	}
}

func (u * PgUserRepo) Create(ctx context.Context, p *models.User) (int64, error)  {
	query := "INSERT INTO public.user(username, email, password_hash,created_at,updated_at) VALUES ($1,$2,$3,$4,$5) returning id;"

	stmt, err := u.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	var lastInsertId int64
	passwordHash, _  := helper.EncodePassword(p.PasswordHash)
	err = u.Conn.QueryRow(query,  p.Username, p.Email, passwordHash,p.CreatedAt,p.UpdatedAt).Scan(&lastInsertId)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return lastInsertId, nil
}

func (u * PgUserRepo) Fetch(ctx context.Context, username string) (*models.User, error)  {
	query := "SELECT * FROM public.user WHERE username = $1"
	user := &models.User{}
	err := u.Conn.Get(user, query, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

