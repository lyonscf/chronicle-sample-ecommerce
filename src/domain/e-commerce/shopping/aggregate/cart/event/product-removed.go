package event

import "github.com/with-hindsight/chronicle/src/domain/core"

type ProductRemoved struct {

	ProductId core.Identifier
}
