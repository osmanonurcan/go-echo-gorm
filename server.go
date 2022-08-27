package main

import (
	"github.com/osmanonurcan/go-test/route"

	"github.com/osmanonurcan/go-test/db"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
