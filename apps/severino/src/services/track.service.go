package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type TrackService struct {
	trackClient interfaces.Client
}

func NewTrackService(container rpc.Container) *TrackService {
	return &TrackService{
		trackClient: container.Elearning.TrackClient,
	}
}

func (this *TrackService) Create(item interface{}) (interface{}, error) {
	return this.trackClient.Create(item)
}

func (this *TrackService) Update(item interface{}) (interface{}, error) {
	return this.trackClient.Update(item)
}

func (this *TrackService) Delete(item interface{}) (interface{}, error) {
	return this.trackClient.Delete(item)
}

func (this *TrackService) ReadOne(item interface{}) (interface{}, error) {
	return this.trackClient.ReadOne(item)
}

func (this *TrackService) ReadAll(search structs.Search) (structs.Result, error) {
	return this.trackClient.ReadAll(search)
}
