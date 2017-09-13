package infrastructure

import (
	"github.com/with-hindsight/chronicle/src/infrastructure/domain/aggregate"
	"github.com/with-hindsight/chronicle/src/infrastructure/domain/player"
	"github.com/with-hindsight/chronicle/src/domain/object"
)

var PlayerRepository = player.NewMemoryRepository()

var AggregateRepository = aggregate.NewMemoryRepository(
	EventLog,
	CommandLog,
)

var DescriptorRepository = object.NewDescriptorRepository(
	DescriptorMap,
)
