package models

import (

)

type Employee struct {
	Number                 uint64 	`json:"customerNumber"`
	FirstName              string 	`json:"customerFirstName"`
	LastName               string 	`json:"customerLastName"`
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