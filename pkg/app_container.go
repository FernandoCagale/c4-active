package pkg

import (
	"github.com/FernandoCagale/c4-active/api/handlers"
	"github.com/FernandoCagale/c4-active/api/routers"
	"github.com/FernandoCagale/c4-active/pkg/domain/active"
	"github.com/google/wire"
)

var Container = wire.NewSet(active.Set, handlers.Set, routers.Set)
