package main

import (
	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/rpc"
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
		AlbumRepository:      repositories.NewAlbumRepository(db),
		TrackRepository:      repositories.NewTrackRepository(db),
		ShelfRepository:      repositories.NewShelfRepository(db),
		PreferenceRepository: repositories.NewPreferenceRepository(db),
		SubjectRepository:    repositories.NewSubjectRepository(db),
	}
	rpc.NewShelfServer(rpcServer.Grpc, repos)
	rpc.NewAlbumServer(rpcServer.Grpc, repos)
	rpc.NewTrackServer(rpcServer.Grpc, repos)
	rpc.NewHomeServer(rpcServer.Grpc, repos)
	rpc.NewPreferenceServer(rpcServer.Grpc, repos)
	rpc.NewSubjectServer(rpcServer.Grpc, repos)

	log.Info("Listening on", port)
	log.Fatal(rpcServer.Start())
}
