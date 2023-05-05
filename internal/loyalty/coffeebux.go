package loyalty

import (
	"context"
	"github.com/google/uuid"
	coffeego "github.com/mamalmaleki/coffeego/internal"
	"github.com/mamalmaleki/coffeego/internal/store"
)

type CoffeeBux struct {
	ID                                    uuid.UUID
	store                                 store.Store
	coffeeLover                           coffeego.CoffeeLover
	FreeDrinksAvailable                   int
	RemainingDrinkPurchasesUntilFreeDrink int
}

func (c *CoffeeBux) AddStamp() {

}

func (c *CoffeeBux) Pay(ctx context.Context, purchases []coffeego.Product) error {
	return nil
}
