package models

import "time"

type User struct {
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user" uri:"id_user" valid:"-"`
	Email_user string     `db:"email" form:"email" json:"email" valid:"email, required~please input in type email"`
	Pass_user  string     `db:"pass" form:"pass" json:"pass" valid:"required~password is required,stringlength(6|1024)~password of at least 6 characters"`
	Role       string     `db:"role" json:"role" form:"role" valid:"-"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
