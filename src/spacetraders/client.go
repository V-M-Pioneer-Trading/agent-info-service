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
	data, err := makeAuthenticatedGetRequest("/my/agent")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed response: %v\n", data)
}

func GetMyShips() {
	data, err := makeAuthenticatedGetRequest("/my/ships")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed response: %v\n", data)
}

func GetMyContracts() {
	data, err := makeAuthenticatedGetRequest("/my/contracts")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed response: %v\n", data)
}

func makeAuthenticatedGetRequest(endpoint string) (interface{}, error) {
	req, err := http.NewRequest("GET", BASE_URL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+BEARER_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
