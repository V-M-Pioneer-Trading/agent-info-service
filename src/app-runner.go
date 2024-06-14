package main

import (
	"net/http"

	"vnm/agent-info-service/api"
	"vnm/agent-info-service/db"
)

func main() {

	db.SetUpDatabase()
	r := api.SetUpRouter()

	http.ListenAndServe(":80", r)

}
