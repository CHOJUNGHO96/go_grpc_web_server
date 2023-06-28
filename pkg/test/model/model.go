package model

type SelectCompanyDepartment struct {
	Id             int    `db:"id"`
	DepartmentName string `db:"department_name"`
	Location       string `db:"location"`
	Manager        string `db:"manager"`
}
