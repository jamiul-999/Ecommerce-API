package handlers

import (
	"ecommerce-api/database"
	"ecommerce-api/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productID")

	pId, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please send a valid product ID", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "Not found", 404)
}
