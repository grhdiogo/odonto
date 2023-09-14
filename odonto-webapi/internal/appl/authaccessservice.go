package appl

// import (
// 	"context"
// 	"time"

// 	"odonto/internal/domain/authaccess"
// 	"odonto/internal/domain/user"
// 	"odonto/internal/infra/config"
// 	"odonto/internal/infra/data/pgclient"
// 	"odonto/internal/infra/fails"
// 	"odonto/internal/infra/utils"

// 	"github.com/robbert229/jwt"
// )

// type authaccessManager interface {
// 	Create(email, secret string) (string, error)
// 	CheckToken(token string) bool
// 	Remove(e *authaccess.Identity) error
// 	IsFirstAccess(ent *authaccess.Entity) bool
// }

// type authaccessManagerImpl struct {
// 	ctx    context.Context
// 	tenant string
// }

// const (
// 	ErrDecodeToken   = "falha ao recuperar dados do token"
// 	ErrEncodeToken   = "falha ao definir dados do token"
// 	ErrValidateToken = "falha ao validar token"
// )

// func (s *authaccessManagerImpl) Create(identifier, secret string) (string, error) {
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return "", fails.ProcessStack(100, ErrGetConnection, err)
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartTx()
// 	if err != nil {
// 		return "", fails.ProcessStack(101, ErrStartTransaction, err)
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new authaccess repository
// 	rep := pgclient.NewAuthAccessRepository(s.ctx, tx, s.tenant)
// 	userRep := pgclient.NewUserRepository(s.ctx, tx)
// 	// find user
// 	userEntity, err := userRep.FindByIdentifierAndSecret(identifier, utils.GetMD5Hash(secret))
// 	if err != nil {
// 		return "", fails.ProcessStack(102, ErrOnFindEntity, err)
// 	}
// 	// verify if is first access based on rg or cpf empty
// 	tkn := utils.GenerateUUID()
// 	expireAt := time.Now().Add(24 * time.Hour)
// 	//create claims
// 	// TODO: ADD DATA
// 	claims := jwt.NewClaim()
// 	claims.Set("tkn", tkn)
// 	claims.Set("rol", userEntity.Kind)
// 	claims.Set("uid", userEntity.ID.PersonPID)
// 	claims.SetTime("exp", expireAt)

// 	//get settings
// 	settings := config.GetSettings()
// 	//create a jwt enconder
// 	jwtSecret := settings.JWTSecret
// 	algorithm := jwt.HmacSha256(jwtSecret)
// 	// enconde token
// 	jwtToken, err := algorithm.Encode(claims)
// 	if err != nil {
// 		return "", fails.ProcessStack(105, ErrEncodeToken, err)
// 	}
// 	if algorithm.Validate(jwtToken) != nil {
// 		return "", fails.ProcessStack(106, ErrValidateToken, err)
// 	}
// 	// create entity
// 	// TODO: ADD DATA
// 	e := &authaccess.Entity{
// 		ID: authaccess.Identity{
// 			Token: tkn,
// 		},
// 		UserUid:   userEntity.ID.PersonPID,
// 		Nick:      "",
// 		ExpiredAt: expireAt,
// 	}
// 	// save access
// 	// store authaccess
// 	err = rep.Store(e)
// 	if err != nil {
// 		return "", fails.ProcessStack(103, ErrOnCreateEntity, err)
// 	}
// 	err = tx.Commit(s.ctx)
// 	if err != nil {
// 		return "", fails.ProcessStack(104, ErrOnCommitTransaction, err)
// 	}
// 	// success
// 	return jwtToken, nil
// }

// func (s *authaccessManagerImpl) CheckToken(token string) bool {
// 	settings := config.GetSettings()
// 	// create a jwt encoder
// 	jwtSecret := settings.JWTSecret
// 	algorithm := jwt.HmacSha256(jwtSecret)
// 	// validate token
// 	err := algorithm.Validate(token)
// 	if err != nil {
// 		return false
// 	}
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return false
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartRTx()
// 	if err != nil {
// 		return false
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new authaccess repository
// 	rep := pgclient.NewAuthAccessRepository(s.ctx, tx, s.tenant)
// 	ent, err := authaccess.NewEntity(token)
// 	if err != nil {
// 		return false
// 	}
// 	// result
// 	_, err = rep.Find(&authaccess.Identity{
// 		Token: ent.ID.Token,
// 	})
// 	if err != nil {
// 		return false
// 	}
// 	err = tx.Commit(s.ctx)
// 	return err == nil
// }

// func (s *authaccessManagerImpl) Remove(identity *authaccess.Identity) error {
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return fails.ProcessStack(112, ErrGetConnection, err)
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartTx()
// 	if err != nil {
// 		return fails.ProcessStack(113, ErrStartTransaction, err)
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new authaccess repository
// 	rep := pgclient.NewAuthAccessRepository(s.ctx, tx, s.tenant)
// 	// delete authaccess
// 	err = rep.Remove(identity)
// 	if err != nil {
// 		return fails.ProcessStack(114, ErrOnRemoveEntity, err)
// 	}
// 	err = tx.Commit(s.ctx)
// 	if err != nil {
// 		return fails.ProcessStack(115, ErrOnCommitTransaction, err)
// 	}
// 	return nil
// }

// func (s *authaccessManagerImpl) IsFirstAccess(ent *authaccess.Entity) bool {
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return false
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartRTx()
// 	if err != nil {
// 		return false
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new repository
// 	userRep := pgclient.NewUserRepository(s.ctx, tx)
// 	// delete authaccess
// 	userEnt, err := userRep.Find(&user.Identity{
// 		PersonPID: ent.UserUid,
// 	})
// 	if err != nil {
// 		return false
// 	}
// 	err = tx.Commit(s.ctx)
// 	if err != nil {
// 		return false
// 	}
// 	// verify
// 	firstAccess := false
// 	// TODO: ADD DATA
// 	if userEnt.ID.PersonPID == "" {
// 		firstAccess = true
// 	}
// 	return firstAccess
// }

// func NewAuthaccessManager(ctx context.Context) authaccessManager {
// 	return &authaccessManagerImpl{
// 		ctx: ctx,
// 	}
// }
