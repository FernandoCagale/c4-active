package main

import (
	"github.com/FernandoCagale/c4-active/api/routers"
	"github.com/FernandoCagale/c4-active/config"
	"github.com/urfave/negroni"
)

func main() {
	env := config.LoadEnv()

	n := negroni.New(
		negroni.NewLogger(),
	)

	router := routers.NewRouter()

	routers.MakeHandlers(router, *n, env)

	n.UseHandler(router)

	n.Run(":8080")
}
