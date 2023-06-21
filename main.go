package main

import (
	"log"
	"net/http"

	"github.com/jagottsicher/myGoWebserver/router"
	"github.com/jagottsicher/myGoWebserver/services"
	"github.com/jagottsicher/myGoWebserver/utils"
)

func main() {
	log.Println("In Main App")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
