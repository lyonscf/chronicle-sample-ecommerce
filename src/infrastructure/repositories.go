package infrastructure

import (
	"github.com/with-hindsight/chronicle/src/application/repository"
	"github.com/with-hindsight/chronicle/src/infrastructure/domain/player"
)

var PlayerRepository = player.NewMemoryRepository()

var AggregateRepository = repository.NewAggregateRepository()

/*
var DescriptorRepository = object.NewDescriptorRepository(
	DescriptorMap,
)
*/
