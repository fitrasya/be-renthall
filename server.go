package main

import (
	"be-renthall/cache"
	"be-renthall/db"
	"be-renthall/route"
)

func main() {
	db.Init()
	cache.Init()
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
