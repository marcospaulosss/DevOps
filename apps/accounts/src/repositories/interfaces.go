package repositories

import "backend/apps/accounts/src/structs"

type Repository interface {
	Create(data interface{}) (interface{}, error)
	ReadAll(search structs.Search) (structs.Result, error)
	ReadOne(data interface{}) (interface{}, error)
	Update(data interface{}) (interface{}, error)
	Delete(data interface{}) (interface{}, error)
}
