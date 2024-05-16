package db

import "products/models"

func GetPage(obj Pagination) ([]models.Product, error) {

	query, err := ReturnQuery(obj)
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(query)
	if err != nil {
		return nil, err
	}

	var AllProduct []models.Product

	for rows.Next() {
		var Product models.Product
		err := rows.Scan(&Product.ID, &Product.Name, &Product.Category, &Product.Price, &Product.Quantity)
		if err != nil {
			return nil, err
		}
		AllProduct = append(AllProduct, Product)
	}
	return AllProduct, nil
}
