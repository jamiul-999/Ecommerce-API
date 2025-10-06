package handlers

import (
	"ecommerce-api/database"
	"ecommerce-api/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please send a valid product ID", 400)
		return
	}

	product := database.Get(pId)
	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}
	util.SendData(w, product, 200)
}
