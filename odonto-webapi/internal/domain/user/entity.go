package user

type UserKind string

const (
	Adm    UserKind = "adm"
	Normal UserKind = "normal"
)

type Entity struct {
	ID         Identity
	Kind       UserKind
	Identifier string
	Secret     string
}
