# Agent Info Service
Service for accessing information about the agent - profile, fleet, contracts, etc.

## Setup and local development

### Required software

* Golang (https://go.dev/doc/install)
* Latest MySql Docker Image (https://hub.docker.com/_/mysql)

### Setting up Environment

**TODO** add more detailed instructions for a fresh project setup

#### Database

To run DB execute the following command:
docker run --name vnm-agent-db -e MYSQL_ROOT_PASSWORD=*your-password* -d mysql:tag

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