package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"vnm/agent-info-service/spacetraders"
)

func SetUpRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/current-agent", func(w http.ResponseWriter, r *http.Request) {

		log.Default().Printf("%s request: to %s", r.Method, r.RequestURI)
		spacetraders.GetMyAgent()
		fmt.Fprintf(w, "You've requested information about the current agent. See server console output for details")
	})

	r.HandleFunc("/agent/{agentId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		agentId := vars["agentId"]

		log.Default().Printf("%s request: to %s", r.Method, r.RequestURI)
		fmt.Fprintf(w, "You've requested information about the agent with id: %s", agentId)
	})

	return r
}
