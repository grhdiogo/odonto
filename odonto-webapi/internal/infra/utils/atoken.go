package utils

import (
	"errors"
	"time"
)

var (
	//ErrEncoding error on encode access to jwt token
	ErrEncodingToken = errors.New("error on encode token")
	//ErrDecodingToken error on decode access to jwt token
	ErrDecodingToken = errors.New("error on decode token")
	//ErrDecodingToken error on get decoded token
	ErrTokenNotDecoded = errors.New("error on get decoded token")
	//ErrInvalidToken error on get decoded token
	ErrInvalidToken = errors.New("error on invalid token")
)

//dto
type AToken struct {
	//extrutura do auth ecosystem
	Kind      string    `yaml:"kind"`
	Nick      string    `json:"nck"` // "nck" : String,             // user nick
	Token     string    `json:"tkn"` // "tkn" : String,             //access token
	UID       string    `json:"uid"` // "uid" : String,             //user id
	ExpiredAt time.Time `json:"exp"` // "exp" : time
	// "rol" : String,             //role
	// "perms": [ { //TODO: definir permissões } ]

}

// func DecodeAuthToken(aToken string) (*AToken, error) {
// 	//get settings
// 	settings := config.GetSettings()
// 	//create a jwt enconder
// 	jwtSecret := settings.JWTSecret
// 	algorithm := jwt.HmacSha256(jwtSecret)
// 	// check validity
// 	if err := algorithm.Validate(aToken); err != nil {
// 		return nil, ErrInvalidToken
// 	}
// 	//extract from atoken
// 	claims, err := algorithm.Decode(aToken)
// 	if err != nil {
// 		return nil, ErrDecodingToken
// 	}
// 	token, err := claims.Get("tkn")
// 	if err != nil {
// 		return nil, ErrTokenNotDecoded
// 	}
// 	nick, err := claims.Get("nck")
// 	if err != nil {
// 		return nil, ErrTokenNotDecoded
// 	}
// 	uid, err := claims.Get("uid")
// 	if err != nil {
// 		return nil, ErrTokenNotDecoded
// 	}
// 	expAt, err := claims.GetTime("exp")
// 	if err != nil {
// 		return nil, ErrTokenNotDecoded
// 	}
// 	return &AToken{
// 		Token:     token.(string),
// 		Nick:      nick.(string),
// 		UID:       uid.(string),
// 		ExpiredAt: expAt,
// 	}, nil
// }

func IsEmptyTime(t time.Time) bool {
	return t.IsZero()
}

//criar estrutura jwtAuthToken do ecosystem
