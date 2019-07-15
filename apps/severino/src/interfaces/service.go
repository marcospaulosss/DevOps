package interfaces

import "backend/apps/severino/src/structs"

type Service interface {
	Create(item interface{}) (interface{}, error)
	Update(item interface{}) (interface{}, error)
	ReadOne(item interface{}) (interface{}, error)
	Delete(item interface{}) (interface{}, error)
	ReadAll(search structs.Search) (structs.Result, error)
}

type AlbumService interface {
	Service
	Publish(item structs.Album) (structs.Album, error)
	Unpublish(item structs.Album) (structs.Album, error)
}
