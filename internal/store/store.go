package store

import (
	"github.com/google/uuid"
	"github.com/mamalmaleki/coffeego"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffeego.Product
}
