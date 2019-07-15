package repositories

import "backend/apps/ecommerce/src/structs"

type Repository interface {
	Create(item interface{}) (interface{}, error)
	ReadAll(search structs.Search) (structs.Result, error)
	ReadOne(item interface{}) (interface{}, error)
	Update(item interface{}) (interface{}, error)
	Delete(item interface{}) (interface{}, error)
}
