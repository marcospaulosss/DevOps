package interfaces

import "backend/apps/severino/src/structs"

type Client interface {
	Create(params interface{}) (interface{}, error)
	ReadAll(search structs.Search) (structs.Result, error)
	ReadOne(params interface{}) (interface{}, error)
	Update(params interface{}) (interface{}, error)
	Delete(params interface{}) (interface{}, error)
}

type AlbumClient interface {
	Client
	Publish(album structs.Album) (structs.Album, error)
	Unpublish(album structs.Album) (structs.Album, error)
}
