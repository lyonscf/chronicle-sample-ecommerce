package infrastructure

import (
	"github.com/with-hindsight/chronicle/src/infrastructure/domain/storage/memory"
	"github.com/with-hindsight/chronicle/src/infrastructure/core/uuid"
)

var IdentifierGenerator = uuid.NewGenerator()

var EventLog = memory.NewLog(IdentifierGenerator)
var CommandLog = memory.NewLog(IdentifierGenerator)
