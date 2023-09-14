package appl

import (
	"context"
	"database/sql"
	"odonto/internal/domain/procedure"
	"odonto/internal/infra/criteria"
	"odonto/internal/infra/data/pgclient"
	"odonto/internal/infra/fails"
	"odonto/internal/infra/utils"
)

type procedureManager interface {
	Create(e *procedure.Entity) error
	CreateAll(e []procedure.Entity) error
	Update(e *procedure.Entity) error
	List(id procedure.Identity, text string, page, limit int) ([]procedure.Entity, error)
	Find(e *procedure.Identity) (*procedure.Entity, error)
	Remove(e *procedure.Identity) error
}

type procedureManagerImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (s *procedureManagerImpl) Create(e *procedure.Entity) error {
	// new procedure repository
	rep := pgclient.NewProcedureRepository(s.ctx, s.tx)
	e.ID.Pid = utils.GenerateUUID()
	// store procedure
	err := rep.Store(e)
	if err != nil {
		return fails.ProcessStack(103, ErrOnCreateEntity, err)
	}
	// success
	return nil
}

func (s *procedureManagerImpl) CreateAll(entities []procedure.Entity) error {
	// new procedure repository
	rep := pgclient.NewProcedureRepository(s.ctx, s.tx)
	// store procedure
	err := rep.StoreAll(entities)
	if err != nil {
		return fails.ProcessStack(105, ErrOnCreateEntity, err)
	}
	// success
	return nil
}

func (s *procedureManagerImpl) Find(identity *procedure.Identity) (*procedure.Entity, error) {
	// new procedure repository
	rep := pgclient.NewProcedureRepository(s.ctx, s.tx)
	// create criteria
	ent, err := rep.Find(*identity)
	if err != nil {
		return nil, fails.ProcessStack(108, ErrOnFindEntity, err)
	}
	return ent, nil
}

func (s *procedureManagerImpl) Update(e *procedure.Entity) error {
	// new procedure repository
	rep := pgclient.NewProcedureRepository(s.ctx, s.tx)
	// update procedure
	err := rep.Update(e)
	if err != nil {
		return fails.ProcessStack(111, ErrOnUpdateEntity, err)
	}
	return nil
}

func (s *procedureManagerImpl) Remove(identity *procedure.Identity) error {
	// new procedure repository
	rep := pgclient.NewProcedureRepository(s.ctx, s.tx)
	// delete procedure
	err := rep.Remove(identity)
	if err != nil {
		return fails.ProcessStack(113, ErrOnRemoveEntity, err)
	}
	return nil
}

func (s *procedureManagerImpl) List(id procedure.Identity, text string, page, limit int) ([]procedure.Entity, error) {
	// new procedure repository
	rep := pgclient.NewProcedureRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.SetPage(page).SetLimit(limit)
	// list procedure
	entities, err := rep.List(builder)
	if err != nil {
		return nil, fails.ProcessStack(116, ErrOnListEntities, err)
	}
	return entities, nil
}

func NewProcedureManager(ctx context.Context, tx *sql.Tx) procedureManager {
	return &procedureManagerImpl{
		ctx: ctx,
		tx:  tx,
	}
}
