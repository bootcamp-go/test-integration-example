package storage

import "fmt"

// NewStorageProductsMap returns a new instance of StorageProductsMap.
func NewStorageProductsMap(db map[int]ProductAttributes, lastId int) *StorageProductsMap {
	return &StorageProductsMap{db: db, lastID: lastId}
}

// StorageProductsMap is an struct that implements StorageProducts interface.
// It stores products in memory using a map.
type StorageProductsMap struct {
	// db is the database of products.
	db map[int]ProductAttributes
	// lastID is the last ID of a product.
	lastID int
}

// Save saves a product.
func (s *StorageProductsMap) Save(product *ProductAttributes) (lastId int, err error) {
	// validate product
	// - check duplicated name
	for _, p := range s.db {
		if p.Name == product.Name {
			err = fmt.Errorf("%w: name", ErrProductDuplicatedField)
			return
		}
	}

	// save product
	s.lastID++
	s.db[s.lastID] = *product

	// get last ID
	lastId = s.lastID

	return
}