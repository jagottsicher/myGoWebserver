package main

import (
	"log"
	"net/http"

	"myGoWebserver/router"
	"myGoWebserver/services"
	"myGoWebserver/utils"
)

func main() {
	log.Println("In Main App")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
