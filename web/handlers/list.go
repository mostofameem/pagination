package handlers

import (
	"net/http"
	"net/url"
	"products/db"
	"products/web/utils"
)

func List(w http.ResponseWriter, r *http.Request) {

	query, err := UrlOperation(r.URL.String())
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err, "Error Loading Query")
		return
	}
	ShowPage(w, query)

}

func UrlOperation(r string) (db.Pagination, error) {

	var qury db.Pagination
	parsedUrl, err := url.Parse(r)
	if err != nil {
		return qury, err
	}
	queryParams := parsedUrl.Query()
	qury.Page = queryParams.Get("page")
	qury.Limit = queryParams.Get("limit")
	qury.Category = queryParams.Get("category")
	qury.Price = queryParams.Get("price_less_then")
	qury.Orderkey = queryParams.Get("orderKey")
	qury.OrderType = queryParams.Get("orderType")

	return qury, nil
}

func ShowPage(w http.ResponseWriter, query db.Pagination) {
	TotalProductchan := make(chan interface{})
	Pagechan := make(chan interface{})

	go db.GetTotalNumberOFProduct(query, TotalProductchan)
	go db.GetPage(query, Pagechan)

	totalProducts := <-TotalProductchan
	page := <-Pagechan
	utils.SendBothData(w, totalProducts, page)

}
