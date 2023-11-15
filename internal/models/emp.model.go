package models

import "time"

type Employee struct {
	Id_Employee   string     `db:"id_employee" form:"id_employee" json:"id_employee" uri:"id_employee" valid:"-"`
	Name_Employee string     `db:"name_employee" form:"name_employee" json:"name_employee" valid:"required~name employee is required"`
	Address       string     `db:"address_employee" form:"address_employee" json:"address_employee" valid:"-"`
	CreatedAt     *time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
