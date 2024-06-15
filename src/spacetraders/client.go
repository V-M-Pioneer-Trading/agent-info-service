package spacetraders

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"vnm/agent-info-service/spacetraders/schema"
)

const BASE_URL = "https://api.spacetraders.io/v2"

// FIXME this is test token untill we have a way to securely transmit a real one
const BEARER_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGlmaWVyIjoiVEVTVC1WTk0tMTQtMDYiLCJ2ZXJzaW9uIjoidjIuMi4wIiwicmVzZXRfZGF0ZSI6IjIwMjQtMDYtMDkiLCJpYXQiOjE3MTgzNjk1OTQsInN1YiI6ImFnZW50LXRva2VuIn0.qYclJL6OrzfoMaM9jMxPqL_q2L-uD_bRNthcwITZANAP5Uod30PiAHj1oPlcy4sY61dIeBvNG-xShwVYaVnBTd1DaMM4hG4ugnBX6JeMLVuR8vybrP2eNLMBs9iiFRcGRumD-PpFH0b7ggHB17PiJEyrwOxkb_rr49-LwAE_nHHNKH8Z9VcqiJ0Rxyz_LMTtdIGlQxQ1qdM4y4J7_fWDpqaa3i-Z8pIJddtgp0ioGf6_frOefHu7Ddps3LEzl1cC5eQp68aQs06jnJIXc4DZp60TAjW08xEBrjxCR5a5No_iwKKezZ6-91DoC3A4dmjcc2zFYCXVrqfvaPCaIkBecQ"

func GetMyAgent() schema.Agent {

	var agent schema.GetMyAgentResponse

	data, err := makeAuthenticatedGetRequest("/my/agent", agent)
	if err != nil {
		panic(err)
	}
	agentResponse := data.(schema.GetMyAgentResponse)
	fmt.Printf("Agent response:\n %+v\n", agentResponse)
	return agentResponse.Data
}

func GetMyShips() []schema.Ship {
	var ships schema.GetMyShipsResponse
	data, err := makeAuthenticatedGetRequest("/my/ships", ships)
	if err != nil {
		panic(err)
	}
	requestResult := data.(schema.GetMyShipsResponse)
	fmt.Printf("Ships response: %v\n", requestResult)
	return requestResult.Data
}

func GetMyContracts() []schema.Contract {
	var contracts schema.GetMyContractsResponse
	data, err := makeAuthenticatedGetRequest("/my/contracts", contracts)
	if err != nil {
		panic(err)
	}
	requestResult := data.(schema.GetMyContractsResponse)
	fmt.Printf("Contracts response: %v\n", requestResult)
	return requestResult.Data
}

func makeAuthenticatedGetRequest(endpoint string, resultType interface{}) (interface{}, error) {
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

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get a successful response from the server")
	}

	// Create a new instance of the resultType using reflection
	resultValue := reflect.New(reflect.TypeOf(resultType)).Interface()

	err = json.NewDecoder(resp.Body).Decode(resultValue)
	if err != nil {
		return nil, err
	}

	// Return the dereferenced value
	return reflect.ValueOf(resultValue).Elem().Interface(), nil
}
