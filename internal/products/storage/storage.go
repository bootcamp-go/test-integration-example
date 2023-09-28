package storage

import "errors"

var (
	// ErrProductDuplicatedField is an error when a product has a duplicated field.
	ErrProductDuplicatedField = errors.New("storage products: product has a duplicated field")
)

// ProductAttributes is an struct with the attributes of a product.
type ProductAttributes struct {
	Name  string
	Type  string
	Count int
	Price float64
}

// StorageProducts is an interface to store products.
type StorageProducts interface {
	// Save saves a product.
	Save(product *ProductAttributes) (lastId int, err error)
}