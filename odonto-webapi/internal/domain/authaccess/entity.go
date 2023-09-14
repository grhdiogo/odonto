package authaccess

import (
	"errors"
	"time"

	"odonto/internal/infra/config"

	"github.com/robbert229/jwt"
)

type UserRole string

type Entity struct {
	Nick      string
	UserUid   string
	Rol       UserRole
	ExpiredAt time.Time
	ID        Identity
}

var (
	//ErrEncodingToken error on encode access to jwt token
	ErrEncodingToken = errors.New("error on encode token")
	//ErrDecodingToken error on decode access to jwt token
	ErrDecodingToken = errors.New("error on decode token")
	//ErrDecodingToken error on get decoded token
	ErrTokenNotDecoded = errors.New("error on get decoded token")
	//ErrInvalidToken error on get decoded token
	ErrInvalidToken = errors.New("invalid token")
)

func NewEntity(jwtToken string) (*Entity, error) {
	// get settings
	settings := config.GetSettings()
	// create a jwt encoder
	jwtSecret := settings.JWTSecret
	algorithm := jwt.HmacSha256(jwtSecret)
	// check validity
	if err := algorithm.Validate(jwtToken); err != nil {
		return nil, ErrInvalidToken
	}
	// extract stoken from jwt token
	claims, err := algorithm.Decode(jwtToken)
	if err != nil {
		return nil, ErrDecodingToken
	}
	uid, err := claims.Get("uid")
	if err != nil {
		return nil, ErrTokenNotDecoded
	}
	tkn, err := claims.Get("tkn")
	if err != nil {
		return nil, ErrTokenNotDecoded
	}
	rol, err := claims.Get("rol")
	if err != nil {
		return nil, ErrTokenNotDecoded
	}
	expAt, err := claims.GetTime("exp")
	if err != nil {
		return nil, ErrTokenNotDecoded
	}
	// create access
	return &Entity{
		ID: Identity{
			Token: tkn.(string),
		},
		UserUid:   uid.(string),
		ExpiredAt: expAt,
		Rol:       UserRole(rol.(string)),
	}, nil
}
