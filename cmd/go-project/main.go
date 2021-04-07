package main

import (
	"context"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "github.com/shahzaibaziz/GOProject"
	"github.com/shahzaibaziz/GOProject/config"
	"github.com/shahzaibaziz/GOProject/gen/restapi"
	"github.com/shahzaibaziz/GOProject/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}
	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	api := handlers.NewHandler(ctx, rt, swaggerSpec)

	server := restapi.NewServer(api)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	server.Host = viper.GetString(config.ServerHost)
	if err != nil {
		panic(err)
	}

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
