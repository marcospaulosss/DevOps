package repositories

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"backend/libs/errors"

	"backend/apps/accounts/src/structs"
	"backend/libs/databases"
	log "backend/libs/logger"
)

type AccountRepository struct {
	db           databases.Database
	queryBuilder AccountQueryBuilder
}

func NewAccountRepository(db databases.Database) AccountRepository {
	return AccountRepository{
		db:           db,
		queryBuilder: NewAccountQueryBuilder(),
	}
}

func (this AccountRepository) Create(item interface{}) (interface{}, error) {
	account := item.(*structs.Account)

	log.Info("Vou adicionar a pre inscricao no BD.")

	query := this.queryBuilder.CreateAccount(account.Type)
	if account.Exists {
		query = this.queryBuilder.CreateAccountLogin(account.Type)
	}

	saved, err := this.upsert(query, *account)
	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (this AccountRepository) ReadOne(item interface{}) (interface{}, error) {
	data := item.(*structs.Account)

	log.Info("Vou validar email/celular e o codigo no BD.", *data)
	query := this.queryBuilder.ValidateAccount(data.Type)

	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Os codigos de validacao nao conferem com o email ou celular fornecidos")
	}

	var account structs.Account
	if err = nstmt.Get(&account, data); err != nil {
		log.Info("Os codigos de validacao nao conferem com o email ou celular fornecidos:", err.Error())
		return nil, status.Errorf(codes.PermissionDenied, "Os codigos de validacao nao conferem com o email ou celular fornecidos")
	}

	log.Info("Encontrei Email/celular e codigo no BD.", account)

	return &account, nil
}

func (this AccountRepository) ReadAll(search structs.Search) (structs.Result, error) {

	return structs.Result{}, nil
}

func (this AccountRepository) Update(data interface{}) (interface{}, error) {

	return nil, nil
}

func (this AccountRepository) Delete(item interface{}) (interface{}, error) {
	account := item.(*structs.Account)

	log.Info("Vou deletar a pre inscricao no BD.", account.ID)

	query := this.queryBuilder.DeleteAccount()

	saved, err := this.upsert(query, *account)
	if err != nil {
		return nil, err
	}

	return saved, nil
}

func (this AccountRepository) upsert(query string, item structs.Account) (*structs.Account, error) {
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no prepare statement.", err)
		return nil, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	var saved structs.Account
	if err = nstmt.Get(&saved, item); err != nil {
		log.Error("Falhei.", item.Email, err)
		return nil, status.Errorf(codes.AlreadyExists, "Usuario ja cadastrado na base")
	}

	log.Info("Ok. Salvei o account:", saved)
	return &saved, nil
}

func (this AccountRepository) FetchCodeExisting(item interface{}) (*structs.Account, error) {
	data := item.(*structs.Account)
	query := this.queryBuilder.SelectAccountByType(data.Type)

	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei na preparacao da query com os parametros passados", query, item)
		return nil, err
	}

	var account structs.Account
	if err = nstmt.Get(&account, item); err != nil {
		log.Info("Nao existe codigos gerados para esse usuario:", err.Error())
		return nil, status.Errorf(codes.NotFound, "Nao existe codigos gerados para esse usuario")
	}

	return &account, nil
}
