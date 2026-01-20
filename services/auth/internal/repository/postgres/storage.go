package postgres

import (
	"context"
	"database/sql"

	"github.com/devvdark0/image-processing-service/services/auth/internal/model"
)

type POSTGRESQLRepository struct {
	db *sql.DB
}

func NewPOSTGRESQLRepository(db *sql.DB) *POSTGRESQLRepository {
	return &POSTGRESQLRepository{
		db: db,
	}
}

func (p *POSTGRESQLRepository) Create(ctx context.Context, email, password string) (int64, error) {
	sql := `INSERT INTO users(email, password) VALUES(?, ?)`

	rs, err := p.db.ExecContext(ctx, sql, email, password)
	if err != nil {
		return 0, err
	}

	userId, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (p *POSTGRESQLRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	sql := `SELECT id, email, password FROM users WHERE email=?`

	var user model.User

	err := p.db.QueryRowContext(ctx, sql, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
