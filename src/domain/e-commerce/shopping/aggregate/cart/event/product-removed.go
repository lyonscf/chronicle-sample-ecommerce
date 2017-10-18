package event

import (
	"github.com/with-hindsight/chronicle/domain"
)

type ProductRemoved struct {

	ProductId domain.Identifier
}
