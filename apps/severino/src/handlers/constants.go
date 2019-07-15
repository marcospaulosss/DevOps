package handlers

const (
	ErrInvalidID      = "O ID fornecido não é um número inteiro e positivo."
	ErrInvalidUUID    = "O ID fornecido não é um UUID válido."
	ErrInvalidJSON    = "O request não é um JSON válido."
	ErrInvalidRequest = "Um ou mais valores enviados não são válidos."
	ErrCreate         = "Falhou ao salvar o item."
	ErrUpdate         = "Falhou ao salvar o item."
	ErrDelete         = "Falhou ao remover o item."
	ErrReadAll        = "Falhou ao buscar os itens."
	ErrReadOne        = "Falhou ao buscar por esse item."
	ErrUnauthorized   = "Falhou usuário não autenticado."
)
