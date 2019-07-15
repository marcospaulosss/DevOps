package rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"backend/apps/ecommerce/src/repositories"
	"backend/apps/ecommerce/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

type ProductServer struct {
	repos repositories.Container
}

func NewProductServer(s *grpc.Server, repos repositories.Container) *ProductServer {
	server := &ProductServer{repos}
	if s != nil {
		pb.RegisterProductServiceServer(s, server)
	}
	return server
}

func (this *ProductServer) ReadOne(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Info("Vou transformar o request protobuf...")
	product := FromPbProduct(in.GetProduct())
	log.Info("Ok. Item:", product.ID)
	result, err := this.repos.ProductRepository.ReadOne(product)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info("Sucesso. Achei o item.")
	saved := result.(structs.Product)
	resp := &pb.ProductResponse{Product: ToPbProduct(saved)}
	log.Info("Vou retornar:", resp.String())
	return resp, nil
}

func (this *ProductServer) Create(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Info("Vou transformar o request:", in.String())
	product := FromPbProduct(in.GetProduct())
	log.Info("Ok. Item:", product.Name)
	result, err := this.repos.ProductRepository.Create(product)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Info("Sucesso. O item foi criado.")
	saved := result.(structs.Product)
	resp := &pb.ProductResponse{
		Product: ToPbProduct(saved),
	}
	log.Info("Vou retornar:", resp.String())
	return resp, nil
}

func (this *ProductServer) ReadAll(ctx context.Context, in *pb.SearchRequest) (*pb.ProductsResponse, error) {
	log.Info("Vou transformar o request protobuf...")
	search := FromPbSearch(in.GetSearch())
	log.Info("Ok. Item:", search.String())
	result, err := this.repos.ProductRepository.ReadAll(search)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	pagination := result.Pagination
	items := result.Items.([]structs.Product)
	log.Info("Sucesso. Achei", len(items), "items.")
	resp := &pb.ProductsResponse{
		Products:   ToPbProducts(items),
		Pagination: ToPbPagination(pagination),
	}
	log.Info("Vou retornar:", resp.String())
	return resp, nil
}

func (this *ProductServer) Update(ctx context.Context, in *pb.ProductRequest) (*pb.ProductResponse, error) {
	return nil, nil
}
