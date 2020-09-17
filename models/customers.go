package models
type Customer struct {
	Number                 uint64 	`json:"customerNumber"`
	Name                   string 	`json:"customerName"`
	FirstName              string 	`json:"customerFirstName"`
	LastName               string 	`json:"customerLastName"`
	Phone                  string 	`json:"phone"`
	AddressLine1           string 	`json:"addressLine1"`
	AddressLine2           string 	`json:"addressLine2"`
	City                   string 	`json:"city"`
	State                  string 	`json:"state"`
	PostCode               string 	`json:"postalCode"`
	Country                string 	`json:"country"`
	// SalesRepEmployeeNumber uint64 	`json:"salesRepEmployeeNumber"`
	// CreditLimit            float64 	`json:"creditLimit"`
}

func FindCustomerByNumber(number uint64) (Customer, error){
	var (
		Number                 uint64
		Name                   string 
		FirstName              string 
		LastName               string 
		Phone                  string 
		AddressLine1           string 
		AddressLine2           NullString 
		City                   string 
		State                  NullString 
		PostCode               NullString 
		Country                string 
	)
	rows, err := db.Query("SELECT customerNumber, customerName, contactFirstName, contactLastName, phone, addressLine1, addressLine2, city, state, postalCode, country FROM customers WHERE customerNumber = ?", number)
	defer rows.Close()
	if err != nil {
		return Customer{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&Number,
			&Name,
			&FirstName,
			&LastName,
			&Phone,
			&AddressLine1,
			&AddressLine2,
			&City,
			&State,
			&PostCode,
			&Country,
		)
	}
	if err != nil {
		return Customer{}, err
	}

	customer := Customer{
		Number: Number,
		Name: Name,
		FirstName: FirstName,
		LastName: LastName,
		Phone: Phone,
		AddressLine1: AddressLine1,
		AddressLine2: AddressLine2.String,
		City: City,
		State: State.String,
		PostCode: PostCode.String,
		Country: Country,
	}

	return customer, err
}
