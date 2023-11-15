package repositories

import (
	"errors"
	"math"
	"strconv"
	"test-hiring/config"
	"test-hiring/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoEmployeeIF interface {
	CreateEmployee(data *models.Employee) (string, error)
	UpdateEmployee(data *models.Employee) (string, error)
	DeleteEmployee(data *models.Employee) (string, error)
	GetDataById(id string) int
	GetInfoEmployee(data *models.Employee) (*config.Result, error)
	GetSumData() int
	GetAllEmployee(data *models.Employee, page string, limit string) (*config.Result, error)
}

type RepoEmployee struct {
	*sqlx.DB
}

func NewEmployee(db *sqlx.DB) *RepoEmployee {
	return &RepoEmployee{db}
}

func (r *RepoEmployee) CreateEmployee(data *models.Employee) (string, error) {
	query := `INSERT INTO public.employee(
		name_employee,
		address_employee
	)VALUES(
		:name_employee,
		:address_employee
	);`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return " ", err
	}

	return "add employee data successfull", nil
}

func (r *RepoEmployee) UpdateEmployee(data *models.Employee) (string, error) {
	query := `UPDATE public.employee SET 
	name_employee=:name_employee, 
	address_employee=:address_employee,
	updated_at=now()
	WHERE id_employee=:id_employee;`
	if data.Name_Employee == "" {
		return "", errors.New("name not be empty")
	}
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update employee data successful", nil
}

func (r *RepoEmployee) DeleteEmployee(data *models.Employee) (string, error) {
	query := `DELETE FROM public.employee WHERE id_employee=:id_employee;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete employee data successful", nil
}

func (r *RepoEmployee) GetDataById(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.employee WHERE id_employee=$1", id)
	return count_data
}

func (r *RepoEmployee) GetInfoEmployee(data *models.Employee) (*config.Result, error) {
	employee_data := []models.Employee{}
	r.Select(&employee_data, `SELECT id_employee, name_employee, address_employee, created_at, updated_at FROM public.employee WHERE id_employee=$1`, data.Id_Employee)
	if len(employee_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: employee_data}, nil
}

func (r *RepoEmployee) GetSumData() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.employee")
	return id
}

func (r *RepoEmployee) GetAllEmployee(data *models.Employee, page string, limit string) (*config.Result, error) {
	emp_data := []models.Employee{}

	var meta_emp config.Metas
	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 5
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := r.GetSumData()

	if count_data <= 0 {
		meta_emp.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_emp.Next = ""
		} else {
			meta_emp.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_emp.Prev = ""
	} else {
		meta_emp.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_emp.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_emp.Last_page = ""
	}

	if count_data != 0 {
		meta_emp.Total_data = strconv.Itoa(count_data)
	} else {
		meta_emp.Total_data = ""
	}
	r.Select(&emp_data, `SELECT id_employee,name_employee, address_employee, created_at, updated_at FROM public.employee LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(emp_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: emp_data, Meta: meta_emp}, nil
}
