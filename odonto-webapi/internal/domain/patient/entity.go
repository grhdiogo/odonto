package patient

import "odonto/internal/domain/person"

type Entity struct {
	ID Identity
	person.Entity
}
