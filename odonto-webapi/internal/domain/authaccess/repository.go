
package authaccess

type Repository interface {
	Find(id *Identity) (*Entity, error)
	Store(e *Entity) error
	Remove(id *Identity) error
}

// ===================================================
// Implementation of identity
// ===================================================

type Identity struct {
	Token string
}
