package active

import "github.com/google/wire"

var ActiveSet = wire.NewSet(NewUseCase, NewRepository, wire.Bind(new(Repository), new(*ActiveRepository)))