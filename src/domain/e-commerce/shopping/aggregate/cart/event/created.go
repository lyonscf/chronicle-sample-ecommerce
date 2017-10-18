package event

import (
	"github.com/with-hindsight/chronicle/domain"
)

type Created struct {

	ShopperId domain.Identifier `json:"shopper_id"`
}
