package db

import (
	"net/url"
)

type Pagination struct {
	Page      string `json:"page"`
	Limit     string `json:"limit"`
	Category  string `json:"category"`
	Price     string `json:"price"`
	Orderkey  string `json:"orderkey"`
	OrderType string `json:"orderType"`
}

func UrlOperation(r string) (Pagination, error) {

	var qury Pagination
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
func ReturnQuery(obj Pagination) (string, error) {

	query := "SELECT * FROM products WHERE 1 = 1"

	if obj.Category != "" {
		query += " AND category = '" + obj.Category + "' "
	}

	if obj.Price != "" {
		query += " and price < "
		query += obj.Price
	}
	obj.Orderkey = "price"
	obj.OrderType = "ASC"

	if obj.Orderkey != "" {
		query += " ORDER BY " + obj.Orderkey + " "
	}
	if obj.OrderType != "" {
		query += obj.OrderType
	}
	PageLimit, Offset := ConfigPageSize(obj)
	query += " LIMIT " + IntToString(PageLimit)
	query += " OFFSET " + IntToString(Offset)
	query += ";"
	return query, nil
}

func ConfigPageSize(obj Pagination) (int, int) {
	PageLimit := 20
	Offset := 0
	if obj.Limit != "" {
		PageLimit = min(StringToInt(obj.Limit), PageLimit)
		Offset = PageLimit * (StringToInt(obj.Page) - 1)
	}
	return PageLimit, Offset
}
func StringToInt(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = (n*10 + int(s[i]-'0'))
	}
	return n
}
func IntToString(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s += string(n%10 + '0')
		n /= 10
	}

	return Reverse(s)
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
