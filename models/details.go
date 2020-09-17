package models

import "time"

type Order struct {
	Number		uint64 	`json:"orderNumber"`
	ItemsNumber uint64	`json:"itemsNumber"`
	TotalPrice 	float64	`json:"totalPrice"`
}

func FindOrdersByCustomerNumber(number uint64) ([]Order, error){
	var (
		Number		uint64 	
		ItemsNumber uint64
		TotalPrice 	float64
	)
	var orders []Order

	rows, err := db.Query("SELECT o.orderNumber, SUM(od.quantityOrdered) AS 'Items number', SUM(od.priceEach * od.quantityOrdered) AS 'Total price' FROM orders o INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber WHERE o.customerNumber = ? GROUP BY o.orderNumber;", number)
	
	defer rows.Close()
	if err != nil {
		return []Order{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&Number,
			&ItemsNumber,
			&TotalPrice,
		)

		order := Order{
			Number: Number,
			ItemsNumber: ItemsNumber,
			TotalPrice: TotalPrice,
		}
		orders = append(orders, order)
	}
	if err != nil {
		return []Order{}, err
	}	

	return orders, err
}