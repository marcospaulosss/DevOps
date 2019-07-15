package testutil

import (
	"context"

	"github.com/stretchr/testify/mock"

	"backend/apps/severino/mocks"
	pb "backend/proto"
)

func MockAlbumServiceClient(method string, ctx context.Context, mockedResult *pb.AlbumResponse, mockedErr error) pb.AlbumServiceClient {
	service := new(mocks.AlbumServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockAlbumServiceClientSearch(method string, ctx context.Context, mockedResult *pb.AlbumsResponse, mockedErr error) pb.AlbumServiceClient {
	service := new(mocks.AlbumServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockProductServiceClient(method string, ctx context.Context, mockedResult *pb.ProductResponse, mockedErr error) pb.ProductServiceClient {
	service := new(mocks.ProductServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockProductServiceClientSearch(method string, ctx context.Context, mockedResult *pb.ProductsResponse, mockedErr error) pb.ProductServiceClient {
	service := new(mocks.ProductServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockShelfServiceClient(method string, ctx context.Context, mockedResult *pb.ShelfResponse, mockedErr error) pb.ShelfServiceClient {
	service := new(mocks.ShelfServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockShelfServiceClientSearch(method string, ctx context.Context, mockedResult *pb.ShelvesResponse, mockedErr error) pb.ShelfServiceClient {
	service := new(mocks.ShelfServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockTrackServiceClient(method string, ctx context.Context, mockedResult *pb.TrackResponse, mockedErr error) pb.TrackServiceClient {
	service := new(mocks.TrackServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockTrackServiceClientSearch(method string, ctx context.Context, mockedResult *pb.TracksResponse, mockedErr error) pb.TrackServiceClient {
	service := new(mocks.TrackServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockAccountServiceClient(method string, ctx context.Context, mockedResult *pb.AccountResponse, mockedErr error) pb.AccountServiceClient {
	service := new(mocks.AccountServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockPreferenceServiceClient(method string, ctx context.Context, mockedResult *pb.PreferenceResponse, mockedErr error) pb.PreferenceServiceClient {
	service := new(mocks.PreferenceServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockSubjectServiceClient(method string, ctx context.Context, mockedResult *pb.SubjectResponse, mockedErr error) pb.SubjectServiceClient {
	service := new(mocks.SubjectServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}

func MockSubjectServiceClientSearch(method string, ctx context.Context, mockedResult *pb.SubjectsResponse, mockedErr error) pb.SubjectServiceClient {
	service := new(mocks.SubjectServiceClient)
	service.On(method, ctx, mock.Anything).Return(mockedResult, mockedErr)
	return service
}
