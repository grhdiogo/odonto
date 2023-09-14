package appointment

import (
	"odonto/internal/infra/criteria"
)

type Repository interface {
	Find(crt criteria.Criteria) (*Entity, error)
	List(crt criteria.Criteria) ([]EntityProxy, int, error)
	Update(e *Entity) error
	Store(e *Entity) error
	Remove(id *Identity) error
}

// ===================================================
// Implementation of identity
// ===================================================

type Identity struct {
	Aid string
}
