package main

import (
	"backend/apps/accounts/libs/notification"
	"backend/apps/accounts/src/repositories"
	"backend/apps/accounts/src/rpc"
	"backend/libs/configuration"
	"backend/libs/databases"
	log "backend/libs/logger"
	"backend/libs/remoteprocedurecall"
)

func main() {
	config := configuration.Get()
	port := ":" + config.GetString("port")

	rpcServer := remoteprocedurecall.NewServer(port)
	if rpcServer == nil {
		log.Fatal("Nao consigo escutar na porta:", port)
	}

	databaseURL := config.GetEnvConfString("database_url")
	if databaseURL == "" {
		log.Fatal("Nao encontrei a URL de conexao do banco de dados.")
	}
	db := databases.NewPostgres(databaseURL)
	db.Connect()

	repos := repositories.Container{
		AccountRepository: repositories.NewAccountRepository(db),
		UserRepository:    repositories.NewUserRepository(db),
		Notifications:     notification.NewNotification(),
	}
	rpc.NewAccountServer(rpcServer.Grpc, repos)
	rpc.NewUserServer(rpcServer.Grpc, repos)

	log.Info("Listening on", port)
	log.Fatal(rpcServer.Start())
}
