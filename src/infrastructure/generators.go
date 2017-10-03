package infrastructure

import (
	uuid_infrastructure "github.com/with-hindsight/chronicle/src/infrastructure/core/uuid"
	context_infrastructure "github.com/with-hindsight/chronicle/src/infrastructure/domain/context"
	"github.com/with-hindsight/chronicle/src/domain/context"
	"github.com/with-hindsight/chronicle/src/domain"
	"github.com/satori/go.uuid"
	"github.com/with-hindsight/chronicle/src/domain/context/aggregate"
)

var IdentifierGenerator = uuid_infrastructure.NewGenerator()

var SampleObjectContext = &context.ObjectContext{
	Id: uuid.FromStringOrNil("08d14870-076e-44ed-a044-6947bdd448da"),
	Name: "create",
	Version: 1,
	AggregateArchetype: aggregate.NewArchetype(
		EventGenerator,
		&context.AggregateContext{
			Id: uuid.FromStringOrNil("f5e5d99a-220a-4946-a0b2-778dac099ff3"),
			Name: "cart",
			Context: &context.Context{
				Id: uuid.FromStringOrNil("958bbea6-89be-4c56-a54e-34a27095941e"),
				Name: "shopping",
				Players: []context.Player{},
				Domain: &domain.Context{
					Id: uuid.FromStringOrNil("b9a185c3-4c80-4c0e-b44a-9e37423beb44"),
					Name: "e-commerce",
				},
			},
		},
	),
}

var ObjectGenerator = context_infrastructure.NewObjectGenerator(
	IdentifierGenerator,
	SampleObjectContext,
)

var CommandGenerator = context.NewCommandGenerator(
	ObjectGenerator,
)

var EventGenerator = context.NewEventGenerator(
	ObjectGenerator,
)
