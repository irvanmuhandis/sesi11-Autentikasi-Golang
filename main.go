package main

import (
	"sesi11/database"
	"sesi11/router"
)

var PORT = ":8080"

func main() {
	database.StartDB()
	router.StartApp().Run(PORT)
}
