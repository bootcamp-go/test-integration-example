package handlers

import (
	"app/internal/products/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for HandlerProducts.Save
func TestHandlerProducts_Save(t *testing.T) {
	t.Run("success to create a product", func(t *testing.T) {
		// assert
		st := storage.NewStorageProductsMap(make(map[int]storage.ProductAttributes), 0)
		hd := NewHandlerProducts(st)

		// act
		req := httptest.NewRequest(
			"POST",
			"/products",
			strings.NewReader(`{"name": "product 1", "type": "type 1", "count": 1, "price": 1.1}`),
		)
		res := httptest.NewRecorder()
		hd.Save(res, req)

		// arrange
		expectedCode := http.StatusCreated
		expectedBody := `{"message":"success to create a product","data":{"id":1,"name":"product 1","type":"type 1","count":1,"price":1.1}}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())		
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("fail to create a product - handler - invalid request body", func(t *testing.T) {
		// assert
		st := storage.NewStorageProductsMap(make(map[int]storage.ProductAttributes), 0)
		hd := NewHandlerProducts(st)

		// act
		req := httptest.NewRequest("POST", "/products", strings.NewReader(`invalid json`))
		res := httptest.NewRecorder()
		hd.Save(res, req)

		// arrange
		expectedCode := http.StatusBadRequest
		expectedBody := `{"message":"invalid request body","data":null}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("fail to create a product - storage - product has a duplicated field", func(t *testing.T) {
		// assert
		db := map[int]storage.ProductAttributes{
			1: {Name: "product 1", Type: "type 1", Count: 1, Price: 1.1},
		}
		st := storage.NewStorageProductsMap(db, 1)
		hd := NewHandlerProducts(st)

		// act
		req := httptest.NewRequest(
			"POST",
			"/products",
			strings.NewReader(`{"name": "product 1", "type": "type 1", "count": 1, "price": 1.1}`),
		)
		res := httptest.NewRecorder()
		hd.Save(res, req)

		// arrange
		expectedCode := http.StatusUnprocessableEntity
		expectedBody := `{"message":"product has a duplicated field","data":null}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}