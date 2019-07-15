package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type SubjectService struct {
	subjectClient interfaces.Client
}

func NewSubjectService(container rpc.Container) *SubjectService {
	return &SubjectService{
		subjectClient: container.Elearning.SubjectClient,
	}
}

func (this *SubjectService) Create(item interface{}) (interface{}, error) {
	return this.subjectClient.Create(item)
}

func (this *SubjectService) Update(item interface{}) (interface{}, error) {
	return this.subjectClient.Update(item)
}

func (this *SubjectService) Delete(item interface{}) (interface{}, error) {
	return this.subjectClient.Delete(item)
}

func (this *SubjectService) ReadOne(item interface{}) (interface{}, error) {
	return this.subjectClient.ReadOne(item)
}

func (this *SubjectService) ReadAll(search structs.Search) (structs.Result, error) {
	return this.subjectClient.ReadAll(search)
}
