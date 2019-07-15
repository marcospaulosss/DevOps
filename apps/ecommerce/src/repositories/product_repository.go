package repositories

import (
	"backend/apps/ecommerce/src/structs"
	"backend/libs/databases"
	log "backend/libs/logger"
)

type ProductRepository struct {
	db           databases.Database
	queryBuilder ProductQueryBuilder
}

func NewProductRepository(db databases.Database) Repository {
	return &ProductRepository{
		db:           db,
		queryBuilder: NewProductQueryBuilder(),
	}
}

func (this *ProductRepository) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Vou buscar produtos...")

	query := this.queryBuilder.ReadAll(search)
	var items []structs.Product

	conn := this.db.GetConnection()
	err := conn.Select(&items, query)
	if err != nil {
		log.Error("Falhei.", err)
		return structs.Result{}, err
	}

	var total int32
	err = conn.Get(&total, this.queryBuilder.Total())
	if err != nil {
		log.Error("Falhei ao obter o total de registros.", err)
		return structs.Result{}, err
	}

	search.Pagination.Total = total
	result := structs.Result{
		Items:      items,
		Pagination: search.Pagination,
	}
	log.Info("Sucesso. Retornando", len(items), "/", total, "items.")
	return result, nil
}

func (this *ProductRepository) ReadOne(item interface{}) (interface{}, error) {
	product := item.(structs.Product)
	log.Info("Vou buscar o produto com ID:", product.ID)

	conn := this.db.GetConnection()
	query := this.queryBuilder.ReadOne()
	result := structs.Product{}
	err := conn.Get(&result, query, product.ID)
	if err != nil {
		log.Error("Falhei ID=", product.ID, err)
		return nil, err
	}
	log.Info("Sucesso. Achei o produto.")
	return result, nil
}

func (this *ProductRepository) Create(item interface{}) (interface{}, error) {
	product := item.(structs.Product)
	log.Info("Vou inserir o produto:", product.Name)
	query := this.queryBuilder.Create()
	return this.upsert(query, product)
}

func (this *ProductRepository) Update(item interface{}) (interface{}, error) {
	product := item.(structs.Product)
	log.Info("Vou atualizar o produto:", product.Name)
	query := this.queryBuilder.Update()
	return this.upsert(query, product)
}

func (this *ProductRepository) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *ProductRepository) upsert(query string, product structs.Product) (interface{}, error) {
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no prepare statement.", err)
		return nil, err
	}

	var saved structs.Product
	if err = nstmt.Get(&saved, product); err != nil {
		log.Error("Falhei.", product.Name, err)
		return nil, err
	}

	log.Info("Sucesso. Salvei o produto:", saved.Name, saved.ID)
	return saved, nil
}
