package main

import (
	"os"
	"sesi11/database"
	"sesi11/router"
)

// var PORT = ":8080"
var PORT = os.Getenv("PORT")

func main() {
	database.StartDB()
	router.StartApp().Run(":" + PORT)
}
