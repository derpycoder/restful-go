package main

import (
	"flag"
	"log"

	"github.com/abhijit-kar/unite-society/restapi"
	"github.com/abhijit-kar/unite-society/restapi/operations"
	"github.com/abhijit-kar/unite-society/restapi/operations/users"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	var portFlag = flag.Int("port", 3000, "Port to run this service on")

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewUniteSocietyAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	api.APIKeyAuth = func(token string) (interface{}, error) {
		if token == "custom-library-164415" {
			return "Access Granted", nil
		}
		return nil, errors.NotImplemented("Invalid ApiKey: " + token)
	}
	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(func(params users.CreateUserParams, principal interface{}) middleware.Responder {
		requestBody := params.Body
		return users.NewCreateUserCreated().WithPayload(requestBody)
	})

	// set the port this service will be run on
	server.Port = *portFlag

	// parse flags
	flag.Parse()

	// TODO: Set Handle

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
