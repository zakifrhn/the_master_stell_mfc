package repositories

import (
	"errors"
	"test-hiring/config"
	"test-hiring/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `INSERT INTO public.users(
		email,
		pass,
		role
	)VALUES(
		:email,
		:pass,
		:role
	);`

	if data.Email_user == " " || data.Pass_user == " " {
		return " ", errors.New("all forms must be filled")
	}

	_, err := r.NamedExec(query, data)
	if err != nil {
		return " ", err
	}

	return " add user data successfull", nil
}

func (r *RepoUser) UpdateUser(data *models.User) (string, error) {
	query := `UPDATE public.users SET 
	email=:email, 
	pass=:pass,
	updated_at=now()
	WHERE id_user=:id_user;`
	if data.Email_user == "" || data.Pass_user == "" {
		return "", errors.New("all forms must be filled")
	}
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update user data successful", nil
}

func (r *RepoUser) DeleteUser(data *models.User) (string, error) {
	query := `DELETE FROM public.users WHERE id_user=:id_user;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete user data successful", nil
}

func (r *RepoUser) GetUserByEmailnId(email string, id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.users WHERE LOWER(email)=LOWER($1) AND id_user!=$2", email, id)
	return count_data
}

func (r *RepoUser) GetUserByEmail(email string) int {
	var count_data int
	r.Get(&count_data, "SELECT count (*) from public.users where LOWER(email)=LOWER($1)", email)
	return count_data
}

func (r *RepoUser) GetUserById(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.users WHERE id_user=$1", id)
	return count_data
}

func (r *RepoUser) GetInfoUserById(data *models.User) (*config.Result, error) {
	users_data := []models.User{}
	r.Select(&users_data, `SELECT id_user, email, "role", created_at, updated_at FROM public.users WHERE id_user=$1`, data.Id_user)
	if len(users_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: users_data}, nil
}
