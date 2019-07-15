package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type AlbumServiceImpl struct {
	albumClient interfaces.AlbumClient
}

func NewAlbumService(container rpc.Container) interfaces.AlbumService {
	return AlbumServiceImpl{
		albumClient: container.Elearning.AlbumClient,
	}
}

func (this AlbumServiceImpl) Create(item interface{}) (interface{}, error) {
	return this.albumClient.Create(item)
}

func (this AlbumServiceImpl) Update(item interface{}) (interface{}, error) {
	return this.albumClient.Update(item)
}

func (this AlbumServiceImpl) Delete(item interface{}) (interface{}, error) {
	return this.albumClient.Delete(item)
}

func (this AlbumServiceImpl) ReadOne(item interface{}) (interface{}, error) {
	return this.albumClient.ReadOne(item)
}

func (this AlbumServiceImpl) ReadAll(search structs.Search) (structs.Result, error) {
	return this.albumClient.ReadAll(search)
}

func (this AlbumServiceImpl) Publish(album structs.Album) (structs.Album, error) {
	return this.albumClient.Publish(album)
}

func (this AlbumServiceImpl) Unpublish(album structs.Album) (structs.Album, error) {
	return this.albumClient.Unpublish(album)
}
