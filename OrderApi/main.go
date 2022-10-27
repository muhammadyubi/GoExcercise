package main

import (
	"excercise.id/orderapi/database"
	_ "excercise.id/orderapi/docs"
	"excercise.id/orderapi/routers"
)

// @title           Order API
// @version         1.0
// @description     API server for orders in "Scalable Webservice with Golang" course from Hacktiv8 × Kominfo.

// @contact.name   Muhammad Al Ayyubi
// @contact.email  M.Yubi@google.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	database.StartDB()
	var port = ":8080"
	routers.StartServer().Run(port)
}
