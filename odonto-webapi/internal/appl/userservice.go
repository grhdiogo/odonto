package appl

// import (
// 	"context"

// 	"odonto/internal/domain/authaccess"
// 	"odonto/internal/domain/user"
// 	"odonto/internal/infra/criteria"
// 	"odonto/internal/infra/data/pgclient"
// 	"odonto/internal/infra/fails"
// 	"odonto/internal/infra/utils"
// )

// type userManager interface {
// 	Create(e *user.Entity) error
// 	Update(e *user.Entity) error
// 	List(id user.Identity, text string, page, limit int) ([]user.Entity, error)
// 	Find(e *user.Identity) (*user.Entity, error)
// 	Remove(e *user.Identity) error
// }

// type userManagerImpl struct {
// 	ctx    context.Context
// 	access *authaccess.Entity
// }

// func (s *userManagerImpl) Create(e *user.Entity) error {
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return fails.ProcessStack(100, ErrGetConnection, err)
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartTx()
// 	if err != nil {
// 		return fails.ProcessStack(101, ErrStartTransaction, err)
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new user repository
// 	rep := pgclient.NewUserRepository(s.ctx, tx)

// 	e.ID.PersonPID = utils.GenerateUUID()
// 	// store user
// 	err = rep.Store(e)
// 	if err != nil {
// 		return fails.ProcessStack(103, ErrOnCreateEntity, err)
// 	}
// 	// success
// 	return tx.Commit(s.ctx)
// }

// func (s *userManagerImpl) Find(identity *user.Identity) (*user.Entity, error) {
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return nil, fails.ProcessStack(106, ErrGetConnection, err)
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartTx()
// 	if err != nil {
// 		return nil, fails.ProcessStack(107, ErrStartTransaction, err)
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new user repository
// 	rep := pgclient.NewUserRepository(s.ctx, tx)
// 	// result
// 	ent, err := rep.Find(identity)
// 	if err != nil {
// 		return nil, fails.ProcessStack(108, ErrOnFindEntity, err)
// 	}
// 	return ent, tx.Commit(s.ctx)
// }

// func (s *userManagerImpl) Update(e *user.Entity) error {
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return fails.ProcessStack(100, ErrGetConnection, err)
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartTx()
// 	if err != nil {
// 		return fails.ProcessStack(110, ErrStartTransaction, err)
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new user repository
// 	rep := pgclient.NewUserRepository(s.ctx, tx)
// 	// update user
// 	err = rep.Update(e)
// 	if err != nil {
// 		return fails.ProcessStack(111, ErrOnUpdateEntity, err)
// 	}
// 	return tx.Commit(s.ctx)
// }

// func (s *userManagerImpl) Remove(identity *user.Identity) error {
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
// 	// new user repository
// 	rep := pgclient.NewUserRepository(s.ctx, tx)
// 	// delete user
// 	err = rep.Remove(identity)
// 	if err != nil {
// 		return fails.ProcessStack(113, ErrOnRemoveEntity, err)
// 	}
// 	return tx.Commit(s.ctx)
// }

// func (s *userManagerImpl) List(id user.Identity, text string, page, limit int) ([]user.Entity, error) {
// 	// get pool of connections
// 	connManager := pgclient.GetManager()
// 	poolConn, err := connManager.GetConn()
// 	if err != nil {
// 		return nil, fails.ProcessStack(114, ErrGetConnection, err)
// 	}
// 	//create transaction repository
// 	tx, err := poolConn.StartTx()
// 	if err != nil {
// 		return nil, fails.ProcessStack(115, ErrStartTransaction, err)
// 	}
// 	defer tx.Rollback(s.ctx)
// 	// new user repository
// 	rep := pgclient.NewUserRepository(s.ctx, tx)
// 	// create criteria
// 	builder := criteria.NewCriteriaBuilder()
// 	builder.SetPage(page).SetLimit(limit)
// 	// list user
// 	entities, err := rep.List(builder)
// 	if err != nil {
// 		return nil, fails.ProcessStack(116, ErrOnListEntities, err)
// 	}
// 	return entities, tx.Commit(s.ctx)
// }

// func NewUserManager(ctx context.Context, access *authaccess.Entity) userManager {
// 	return &userManagerImpl{
// 		ctx:    ctx,
// 		access: access,
// 	}
// }
