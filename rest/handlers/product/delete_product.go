package product

import (
	"ecommerce-api/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	err = h.svc.Delete(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	}

	util.SendData(w, http.StatusOK, "Successfully deleted product")
}
