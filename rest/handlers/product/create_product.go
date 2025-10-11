package product

import (
	"ecommerce-api/database"
	"ecommerce-api/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please provide a valid json", 400)
		return
	}

	createdProudct := database.Store(newProduct)
	util.SendData(w, createdProudct, 201)
}
