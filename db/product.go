package db

import (
	"log"
	"products/models"
)

type Pageinfo struct {
	Page      string `json:"page"`
	Limit     string `json:"limit"`
	TotalPage string `json:"totalpage"`
	Totaldata string `josn:"totaldata"`
}

func GetPage(obj Pagination, PageChan chan interface{}) {

	query, err := ReturnQuery(obj)
	if err != nil {
		PageChan <- 0
	}

	rows, err := Db.Query(query)
	if err != nil {
		PageChan <- 0
	}

	var AllProduct []models.Product

	for rows.Next() {
		var Product models.Product
		err := rows.Scan(&Product.ID, &Product.Name, &Product.Category, &Product.Price, &Product.Quantity)
		if err != nil {
			PageChan <- 0
		}
		AllProduct = append(AllProduct, Product)
	}
	PageChan <- AllProduct
}
func GetTotalNumberOFProduct(obj Pagination, total chan interface{}) {

	defer func() {
		total <- ""
	}()

	query, err := TotalProductQuery(obj)
	if err != nil {
		log.Println("Error building total product query:", err)
		return
	}

	var totalProducts int

	err = Db.QueryRow(query).Scan(&totalProducts)
	if err != nil {
		log.Println(query)
		log.Println("Error executing total product query:", err)
		return
	}
	var Pginfo Pageinfo

	Pginfo.Page = obj.Page
	Pginfo.Limit = obj.Limit
	Pginfo.TotalPage = TotalPage(obj, totalProducts)
	Pginfo.Totaldata = IntToString(totalProducts)

	total <- Pginfo
}
func TotalPage(obj Pagination, totalProducts int) string {
	TotalNumberOfPageINT := totalProducts / StringToInt(obj.Limit)
	if totalProducts%StringToInt(obj.Limit) != 0 {
		TotalNumberOfPageINT++
	}
	return IntToString(TotalNumberOfPageINT)
}
