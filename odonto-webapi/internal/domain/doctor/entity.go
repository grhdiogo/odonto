package doctor

import "odonto/internal/domain/person"

type Entity struct {
	person.Entity
	ID Identity
}
