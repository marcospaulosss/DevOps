package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	app "backend/apps/severino/src"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/services"
	"backend/libs/configuration"
)

func getRpcConn(url string) *grpc.ClientConn {
	client := rpc.NewRpcClient(url)
	return client.Connect()
}

func main() {
	config := configuration.Get()

	elearningConn := getRpcConn(config.GetEnvConfString("connections.elearning"))
	accountsConn := getRpcConn(config.GetEnvConfString("connections.accounts"))
	ecommerceConn := getRpcConn(config.GetEnvConfString("connections.ecommerce"))
	connections := []*grpc.ClientConn{elearningConn, accountsConn, ecommerceConn}
	app.SetRpcConnections(connections)

	rpcContainer := rpc.Container{
		Elearning: rpc.NewElearningRpcContainer(elearningConn),
		Accounts:  rpc.NewAccountsRpcContainer(accountsConn),
		Ecommerce: rpc.NewEcommerceRpcContainer(ecommerceConn),
	}

	serviceContainer := services.Container{
		Album:      services.NewAlbumService(rpcContainer),
		Shelf:      services.NewShelfService(rpcContainer),
		Track:      services.NewTrackService(rpcContainer),
		Home:       services.NewHomeService(rpcContainer),
		User:       services.NewUserService(rpcContainer),
		Account:    services.NewAccountService(rpcContainer),
		Product:    services.NewProductService(rpcContainer),
		Preference: services.NewPreferenceService(rpcContainer),
		Subject:    services.NewSubjectService(rpcContainer),
	}

	var application *app.Application
	application = app.New(serviceContainer)
	server := application.GetServer()
	port := ":" + config.GetString("port")
	fmt.Println("Listening on", port)
	log.Fatal(server.Start(port))
}
