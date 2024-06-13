package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"vnm/agent-info-service/db"
)

func main() {
	r := mux.NewRouter()

	db.GetDatabase()

	r.HandleFunc("/current-agent", func(w http.ResponseWriter, r *http.Request) {

		log.Default().Printf("%s request: to %s", r.Method, r.RequestURI)
		fmt.Fprintf(w, "You've requested information about the current agent")
	})

	r.HandleFunc("/agent/{agentId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		agentId := vars["agentId"]

		log.Default().Printf("%s request: to %s", r.Method, r.RequestURI)
		fmt.Fprintf(w, "You've requested information about the agent with id: %s", agentId)
	})

	http.ListenAndServe(":80", r)

}
