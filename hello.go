package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/current-agent", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "You've requested information about the current agent")
	})

	r.HandleFunc("/info/{agentId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		agentId := vars["agentId"]

		fmt.Fprintf(w, "You've requested information about the agent with id: %s", agentId)
	})

	http.ListenAndServe(":80", r)
}
