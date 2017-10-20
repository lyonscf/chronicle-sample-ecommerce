package active_carts

import (
	"github.com/with-hindsight/chronicle/domain"
	"errors"
)

var (
	ErrShopperHasActiveCart = errors.New("This shopper already has an active cart.")
	ErrCartDoesNotExists = errors.New("The cart cannot be removed, because it does not exist.")
)

type Projection interface {

	Add(cart domain.Identifier, shopper domain.Identifier) error

	Remove(cart domain.Identifier) error
}
