package user

import "odonto/internal/infra/criteria"

type Repository interface {
	Find(id *Identity) (*Entity, error)
	List(criteria criteria.Criteria) ([]Entity, error)
	Update(e *Entity) error
	Store(e *Entity) error
	Remove(id *Identity) error
	FindByIdentifierAndSecret(i, s string) (*Entity, error)
}

// ===================================================
// Implementation of identity
// ===================================================

type Identity struct {
	PersonPID string
}
