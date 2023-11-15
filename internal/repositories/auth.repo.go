package repositories

import (
	"errors"
	"test-hiring/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_AuthIF interface {
	GetCountByEmail(email string) int
	GetUser(data *models.Auth) (*models.User, error)
}

type Repo_Auth struct {
	*sqlx.DB
}

func New_Auth(db *sqlx.DB) *Repo_Auth {
	return &Repo_Auth{db}
}

func (r *Repo_Auth) GetCountByEmail(email string) int {
	var count_data int
	r.Get(&count_data, "SELECT count (*) FROM public.users WHERE LOWER(email)=LOWER($1)", email)
	return count_data
}

func (r *Repo_Auth) GetUser(data *models.Auth) (*models.User, error) {
	var result models.User

	q := `SELECT id_user, email, "role", pass FROM public.users WHERE email = ?`

	if err := r.Get(&result, r.Rebind(q), data.Email_user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("email not found")
		}

		return nil, err
	}

	return &result, nil
}
