package handlers

import (
	"net/http"
	"products/db"
	"products/web/utils"
)

func List(w http.ResponseWriter, r *http.Request) {

	query, err := db.UrlOperation(r.URL.String())
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err, "Error Loading Query")
		return
	}

	page, err := db.GetPage(query)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err, "Error Loading Products")
		return
	}

	utils.SendData(w, page)
}
