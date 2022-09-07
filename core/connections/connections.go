package connections

import "go.uber.org/fx"

var Module = fx.Provide(
	GetDatabaseConnection,
	getRedisConnection,
)
