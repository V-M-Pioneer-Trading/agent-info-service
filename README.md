# Agent Info Service
Service for accessing information about the agent - profile, fleet, contracts, etc.

## Setup and local development

### Required software

* Golang (https://go.dev/doc/install)
* Latest MySql Docker Image (https://hub.docker.com/_/mysql)

### Running application locally

**TODO** add more detailed instructions for a fresh project setup

To run DB execute the following command in the root directory:
> docker compose up

To run application, switch to "/src" directory and execute:
> go run .

You can test the applcation by navigating to http://localhost/current-agent

## References

### Basic Golang web docs
https://gowebexamples.com/

#### Spacetraders API reference

https://spacetraders.stoplight.io/docs/spacetraders/11f2735b75b02-space-traders-api

##### GET current agent
https://api.spacetraders.io/v2/my/agent

##### GET my ships
https://api.spacetraders.io/v2/my/ships

##### GET my contracts
https://api.spacetraders.io/v2/my/contracts