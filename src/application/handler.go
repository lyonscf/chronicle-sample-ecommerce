package application

import (
	"github.com/with-hindsight/chronicle/src/domain/handler"
	"github.com/with-hindsight/chronicle/src/domain/command"
	"github.com/lyonscf/chronicle-sample-ecommerce/src/infrastructure"
	"github.com/with-hindsight/chronicle/src/domain/event"
	"github.com/with-hindsight/chronicle/src/domain/core"
	"github.com/with-hindsight/chronicle/src/domain/object"
)

var IdentifierGenerator = infrastructure.IdentifierGenerator

var ObjectGenerator = object.NewGenerator(
	IdentifierGenerator,
	infrastructure.DescriptorRepository,
)

var CommandHandler = handler.NewCommandHandler(
	infrastructure.AggregateRepository,
	infrastructure.PlayerRepository,
)

var CommandGenerator = command.NewGenerator(
	ObjectGenerator,
)

func Handle(identifier core.Identifier, payload core.Payload) (events []*event.Event, err error) {
	return CommandHandler.Handle(
		CommandGenerator.New(
			identifier,
			payload,
		),
	)
}
