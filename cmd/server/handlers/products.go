package handlers

import (
	"app/internal/products/storage"
	"encoding/json"
	"errors"
	"net/http"
)

// NewHandlerProducts returns a new instance of HandlerProducts.
func NewHandlerProducts(st storage.StorageProducts) *HandlerProducts {
	return &HandlerProducts{st: st}
}

// HandlerProducts is an struct with methods to handle requests (handlers)
type HandlerProducts struct {
	// st is the storage of products.
	st storage.StorageProducts
}

// Save handles the request to save a product.

type RequestBodySave struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int	  `json:"count"`
	Price float64 `json:"price"`
}
type ProductJSON struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Count int	  `json:"count"`
	Price float64 `json:"price"`
}
func (h *HandlerProducts) Save(w http.ResponseWriter, r *http.Request) {
	// request
	var req RequestBodySave
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		code := http.StatusBadRequest
		body := map[string]any{"message": "invalid request body", "data": nil}

		w.WriteHeader(code); w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	// process
	// - deserialize
	product := storage.ProductAttributes{
		Name:  req.Name,
		Type:  req.Type,
		Count: req.Count,
		Price: req.Price,
	}
	// - save
	lastId, err := h.st.Save(&product)
	if err != nil {
		var code int; var body map[string]any
		switch {
		case errors.Is(err, storage.ErrProductDuplicatedField):
			code = http.StatusUnprocessableEntity
			body = map[string]any{"message": "product has a duplicated field", "data": nil}
		default:
			code = http.StatusInternalServerError
			body = map[string]any{"message": "internal server error", "data": nil}
		}

		w.WriteHeader(code); w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}

	// response
	code := http.StatusCreated
	body := map[string]any{
		"message": "success to create a product",
		"data": ProductJSON{Id: lastId, Name: product.Name, Type: product.Type, Count: product.Count, Price: product.Price},
	}

	w.WriteHeader(code); w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}
