package handlers

import "github.com/google/wire"

var ApplicationHandlersSet = wire.NewSet(NewHealth, NewActive)
