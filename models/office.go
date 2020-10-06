package models

type Office struct {
	Code		uint64 	`json:"officeCode"`
	City		string	`json:"city"`	
	Phone		string	`json:"phone"`
	Address1		string	`json:"adressLine1"`
	Address2		string	`json:"adressLine2"`
	State		string	`json:"state"`
	Country		string	`json:"country"`
	PostalCode	string	`json:"postalCode"`
	Territory 	string	`json:"territory"`	
}

func FindOfficeByCode(code uint64) (Office, error){
	var (
		Code		uint64 	
		City, Phone, Address1, Address2, State, Country, PostalCode, Territory string	
	)
	row := db.QueryRow("SELECT o.officeCode, o.city, o.phone, o.addressLine1, o.addressLine2, o.state, o.country, o.postalCode, o.territory FROM offices o WHERE o.officeCode = ?;", code)

	err := row.Scan(
		&Code,
		&City,
		&Phone,
		&Address1,
		&Address2,
		&State,
		&Country,
		&PostalCode,
		&Territory,
	)

	if err != nil {
		return Office{}, err
	}

	office := Office{
		Code: Code,
		City: City,
		Phone: Phone,
		Address1: Address1,
		Address2: Address2,
		State: State,
		Country: Country,
		PostalCode: PostalCode,
		Territory: Territory,
	}

	return office, err
}