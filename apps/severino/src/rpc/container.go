package rpc

import (
	"context"

	"google.golang.org/grpc"

	"backend/apps/severino/src/interfaces"
	pb "backend/proto"
)

type Container struct {
	Elearning ElearningRpcContainer
	Accounts  AccountsRpcContainer
	Ecommerce EcommerceRpcContainer
}

type ElearningRpcContainer struct {
	AlbumClient      interfaces.AlbumClient
	ShelfClient      interfaces.Client
	TrackClient      interfaces.Client
	HomeClient       interfaces.Client
	PreferenceClient interfaces.Client
	SubjectClient    interfaces.Client
}

type AccountsRpcContainer struct {
	UserClient    interfaces.Client
	AccountClient interfaces.Client
}

type EcommerceRpcContainer struct {
	ProductClient interfaces.Client
}

func NewElearningRpcContainer(conn *grpc.ClientConn) ElearningRpcContainer {
	ctx := context.Background()
	return ElearningRpcContainer{
		AlbumClient:      NewAlbumClient(pb.NewAlbumServiceClient(conn), ctx),
		TrackClient:      NewTrackClient(pb.NewTrackServiceClient(conn), ctx),
		ShelfClient:      NewShelfClient(pb.NewShelfServiceClient(conn), ctx),
		HomeClient:       NewHomeClient(pb.NewHomeServiceClient(conn), ctx),
		PreferenceClient: NewPreferenceClient(pb.NewPreferenceServiceClient(conn), ctx),
		SubjectClient:    NewSubjectClient(pb.NewSubjectServiceClient(conn), ctx),
	}
}

func NewAccountsRpcContainer(conn *grpc.ClientConn) AccountsRpcContainer {
	ctx := context.Background()
	return AccountsRpcContainer{
		UserClient:    NewUserClient(conn, ctx),
		AccountClient: NewAccountClient(pb.NewAccountServiceClient(conn), ctx),
	}
}

func NewEcommerceRpcContainer(conn *grpc.ClientConn) EcommerceRpcContainer {
	ctx := context.Background()
	return EcommerceRpcContainer{
		//ProductClient: NewProductClient(conn, ctx),
		ProductClient: NewProductClient(pb.NewProductServiceClient(conn), ctx),
	}
}
