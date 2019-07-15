package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type PreferenceService struct {
	preferenceClient interfaces.Client
}

func NewPreferenceService(container rpc.Container) *PreferenceService {
	return &PreferenceService{
		preferenceClient: container.Elearning.PreferenceClient,
	}
}

func (this *PreferenceService) Create(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *PreferenceService) Update(item interface{}) (interface{}, error) {
	return this.preferenceClient.Update(item)
}

func (this *PreferenceService) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *PreferenceService) ReadOne(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *PreferenceService) ReadAll(search structs.Search) (structs.Result, error) {
	return structs.Result{}, nil
}
