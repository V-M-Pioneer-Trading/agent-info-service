package spacetraders

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_URL = "https://api.spacetraders.io/v2"

// FIXME this is test token untill we have a way to securely transmit a real one
const BEARER_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiVEVTVC1WTk0tMTQtMDYiLCJ2ZXJzaW9uIjoidjIuMi4wIiwicmVzZXRfZGF0ZSI6IjIwMjQtMDYtMDkiLCJpYXQiOjE3MTgzNjk1OTQsInN1YiI6ImFnZW50LXRva2VuIn0.qYclJL6OrzfoMaM9jMxPqL_q2L-uD_bRNthcwITZANAP5Uod30PiAHj1oPlcy4sY61dIeBvNG-xShwVYaVnBTd1DaMM4hG4ugnBX6JeMLVuR8vybrP2eNLMBs9iiFRcGRumD-PpFH0b7ggHB17PiJEyrwOxkb_rr49-LwAE_nHHNKH8Z9VcqiJ0Rxyz_LMTtdIGlQxQ1qdM4y4J7_fWDpqaa3i-Z8pIJddtgp0ioGf6_frOefHu7Ddps3LEzl1cC5eQp68aQs06jnJIXc4DZp60TAjW08xEBrjxCR5a5No_iwKKezZ6-91DoC3A4dmjcc2zFYCXVrqfvaPCaIkBecQ"

func GetMyAgent() {

	req, err := http.NewRequest("GET", BASE_URL+"/my/agent", nil)
	if err != nil {
		// todo handle error
		panic(err)
	}

	// Set the Authorization header with the bearer token
	req.Header.Set("Authorization", "Bearer "+BEARER_TOKEN)

	// Use http.DefaultClient.Do to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// todo handle error
		panic(err)
	}
	defer resp.Body.Close()

	// todo create a schema for this json
	// Decode the response body into a JSON object
	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed response: %v", j)
}

func GetMyShips() {

}

func getMyContracts() {

}
