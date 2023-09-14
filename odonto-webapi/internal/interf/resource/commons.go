package resource
import (
	"context"
	"encoding/json"
	"net/http"

)

// ==============================================================
// Constants
// ==============================================================

const (
	ErrOnParseRequest = "dados da requisição inválidos"
	ErrOnCreateEntity = "falha ao criar entidade"
	ErrOnFindEntity   = "falha ao buscar entidade"
	ErrOnListEntities = "falha ao listar entidades"
	ErrOnRemoveEntity = "falha ao deletar entidade"
	ErrOnUpdateEntity = "falha ao atualizar entidade"
	ErrNotAuthorized  = "usuário não autorizado para esta ação"
	ErrOnSigning      = "falha ao assinar documento"
	ErrOnCancel       = "falha ao cancelar documento"
	ErrOnDecodeToken  = "falha ao decodifcar token"
)

// ==============================================================
// Response json
// ==============================================================

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}