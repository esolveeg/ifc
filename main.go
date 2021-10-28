package main

import (
	"ifc/db"
	"ifc/handler"
	"ifc/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")
	db.InitDatabase()
	db := db.DBConn

	h := handler.NewHandler(db)
	h.Register(v1)
	r.Logger.Fatal(r.Start(":8585"))
}
