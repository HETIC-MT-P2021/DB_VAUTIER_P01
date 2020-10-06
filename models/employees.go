package models

import (

)

type Employee struct {
	Number                 uint64 	`json:"employeeNumber"`
	FirstName              string 	`json:"FirstName"`
	LastName               string 	`json:"LastName"`
}

type EmployeeList struct {
	Number      uint64 	`json:"employeeNumber"`
	FirstName   string 	`json:"FirstName"`
	LastName    string 	`json:"LastName"`	
	JobTitle    string 	`json:"jobTitle"`	
	Email		string 	`json:"email"`
	OfficeCode	string 	`json:"officeCode"`
	City		string 	`json:"city"`
	Address1	string 	`json:"addressLine1"`
}

func FindEmployeesByOfficeCode(code uint64) ([]Employee, error){
	var (
		Number		uint64 	
		FirstName, LastName string	
	)
	var employees []Employee

	rows, err := db.Query("SELECT e.employeeNumber, e.firstName, e.lastName FROM employees e INNER JOIN offices o ON e.officeCode = o.officeCode WHERE o.officeCode = ?;", code)
	
	defer rows.Close()
	if err != nil {
		return []Employee{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&Number,
			&FirstName,
			&LastName,
		)

		employee := Employee{
			Number: Number,
			FirstName: FirstName,
			LastName: LastName,
		}
		employees = append(employees, employee)
	}
	if err != nil {
		return []Employee{}, err
	}	

	return employees, err
}

func FindEmployees() ([]EmployeeList, error){
	var (
		Number uint64
		FirstName, LastName, JobTitle, Email, OfficeCode, City, Address1 string
	)
	var employees []EmployeeList

	rows, err := db.Query("SELECT e.employeeNumber, e.firstName, e.lastName, e.jobTitle, e.email, e.officeCode, o.city, o.addressLine1 FROM employees e INNER JOIN offices o ON o.officeCode = e.officeCode ORDER BY e.officeCode, e.lastName")
	
	defer rows.Close()
	if err != nil {
		return []EmployeeList{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&Number,
			&FirstName,
			&LastName,
			&JobTitle,
			&Email,
			&OfficeCode,
			&City,
			&Address1,
		)

		employee := EmployeeList{
			Number: Number,
			FirstName: FirstName,
			LastName: LastName,
			JobTitle: JobTitle,
			Email: Email,
			OfficeCode: OfficeCode,
			City: City,
			Address1: Address1,
		}
		employees = append(employees, employee)
	}
	if err != nil {
		return []EmployeeList{}, err
	}	

	return employees, err
}