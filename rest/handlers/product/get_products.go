package product

import (
	"ecommerce-api/domain"
	"ecommerce-api/util"
	"net/http"
	"strconv"
)

type Pagination struct {
	Data       []*domain.Product `json:"items"`
	Page       int64             `json:"page"`
	Limit      int64             `json:"limit"`
	TotalItems int64             `json:"totalItems"`
	TotalPages int64             `json:"totalPages"`
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	// get page from query params
	// get limit from query params
	reqQuery := r.URL.Query()
	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	cnt, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	paginatedData := Pagination{
		Data:       productList,
		Page:       page,
		Limit:      limit,
		TotalItems: cnt,
		TotalPages: cnt / limit,
	}

	util.SendData(w, http.StatusOK, paginatedData)
}
