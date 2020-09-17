package models

type OrderDetails struct {
	Productcode      	string  `json:"productcode"`
	ProductName 		string  `json:"productName"`
	ProductDescription  string `json:"productDescription"`
	Quantity  			uint64 `json:"quantity"`
	UnitPrice  			float64 `json:"unitPrice"`
	LinePrice			float64	`json:"linePrice"`
}

func FindOrderDetailsByOrderNumber(number uint64) ([]OrderDetails, error) {
	var (
		Productcode      	string
		ProductName 		string
		ProductDescription  string
		Quantity  			uint64
		UnitPrice  			float64
		LinePrice			float64
	)
	var orderDetails []OrderDetails

	rows, err := db.Query("SELECT od.productCode,p.productName, p.productDescription, od.quantityOrdered, od.priceEach, (od.quantityOrdered*od.priceEach) FROM orders o INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber INNER JOIN products p ON od.productCode = p.productCode WHERE o.orderNumber = ? ORDER BY od.orderLineNumber", number)

	defer rows.Close()
	if err != nil {
		return []OrderDetails{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&Productcode,
			&ProductName,
			&ProductDescription,
			&Quantity,
			&UnitPrice,
			&LinePrice,
		)

		orderDetail := OrderDetails{
			Productcode: Productcode,
			ProductName: ProductName,
			ProductDescription: ProductDescription,
			Quantity: Quantity,
			UnitPrice: UnitPrice,
			LinePrice: LinePrice,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	if err != nil {
		return []OrderDetails{}, err
	}

	return orderDetails, err
}
