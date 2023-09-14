package person

import (
	"odonto/internal/infra/criteria"
)

type Repository interface {
	Find(crt criteria.Criteria) (*Entity, error)
	List(crt criteria.Criteria) ([]Entity, error)
	Update(e *Entity) error
	Store(e *Entity) error
	StoreAll(e []Entity) error
	Remove(id *Identity) error
}

// ===================================================
// Implementation of identity
// ===================================================

type Identity struct {
	Pid string     
}
