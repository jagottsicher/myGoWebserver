package main

import (
	"net/http"
	"github.com/YashaswiNayak99/gorilla-pq-test/utils"
	"github.com/YashaswiNayak99/gorilla-pq-test/services"
	"github.com/YashaswiNayak99/gorilla-pq-test/router"
	"log"
)

func main() {
	log.Println("In Main App")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()
	
	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
