package rpc

import (
	"context"

	"github.com/labstack/gommon/log"

	"backend/apps/severino/src/adapters"
	"backend/libs/errors"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/structs"
	pb "backend/proto"
)

type PreferenceClientImpl struct {
	service pb.PreferenceServiceClient
	ctx     context.Context
}

func NewPreferenceClient(service pb.PreferenceServiceClient, ctx context.Context) interfaces.Client {
	return &PreferenceClientImpl{service, ctx}
}

func (this *PreferenceClientImpl) Create(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *PreferenceClientImpl) Update(item interface{}) (interface{}, error) {
	preference := item.(structs.Preference)
	req := adapters.CreatePreferenceRequest(preference)
	_, err := this.service.Update(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	return preference, nil
}

func (this *PreferenceClientImpl) ReadOne(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *PreferenceClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	return structs.Result{}, nil
}

func (this *PreferenceClientImpl) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}
