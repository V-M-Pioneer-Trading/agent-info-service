package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"vnm/agent-info-service/spacetraders"
	"vnm/agent-info-service/spacetraders/schema"
)

type CurrentAgentResponse struct {
	Agent     schema.Agent      `json:"agent"`
	Ships     []schema.Ship     `json:"ships"`
	Contracts []schema.Contract `json:"contracts"`
}

func SetUpRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/current-agent", func(w http.ResponseWriter, r *http.Request) {

		log.Default().Printf("%s request: to %s", r.Method, r.RequestURI)

		// Create a CurrentAgentResponse instance
		var response CurrentAgentResponse
		response.Agent = spacetraders.GetMyAgent()
		response.Ships = spacetraders.GetMyShips()
		response.Contracts = spacetraders.GetMyContracts()

		// Marshal the response into JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response
		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.HandleFunc("/agent/{agentId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		agentId := vars["agentId"]

		log.Default().Printf("%s request: to %s", r.Method, r.RequestURI)
		fmt.Fprintf(w, "You've requested information about the agent with id: %s", agentId)
	})

	return r
}
