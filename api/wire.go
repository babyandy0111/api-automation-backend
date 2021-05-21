//+build wireinject

package api

import (
	"github.com/google/wire"
)

var OrmSet = wire.NewSet(InitXorm)
var CacheSet = wire.NewSet(InitRedis)
