package person

import (
	"time"
)

type Entity struct { 
	Name string  
	Cpf string  
	Email string  
	Birthdate time.Time  
	ID Identity
}
