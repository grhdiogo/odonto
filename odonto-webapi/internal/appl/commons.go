package appl

import (
	"errors"
	"fmt"
	"strings"
)

const (
	ErrGetConnection       = "falha na conexão com o banco de dados"
	ErrStartTransaction    = "falha ao iniciar transação"
	ErrOnCreateEntity      = "falha ao criar entidade"
	ErrOnFindEntity        = "falha ao buscar entidade"
	ErrOnUpdateEntity      = "falha ao buscar entidade"
	ErrOnListEntities      = "falha ao listar entidades"
	ErrOnRemoveEntity      = "falha ao deletar entidade"
	ErrOnCommitTransaction = "falha ao commitar transação"
)

func DefaultSizeValidation(field, filedName string, min, max int) error {
	errs := []string{}
	msgMin := `o campo %s deve ter no mínimo %d caractere(s)`
	msgMax := `o campo %s deve ter no máximo %d caractere(s)`
	if min > 0 && len(field) < min {
		errs = append(errs, fmt.Sprintf(msgMin, filedName, min))
	}
	if len(field) > max {
		errs = append(errs, fmt.Sprintf(msgMax, filedName, min))
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ","))
	}
	return nil
}
