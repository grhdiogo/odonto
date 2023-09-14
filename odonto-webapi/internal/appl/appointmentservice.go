package appl

import (
	"context"
	"database/sql"
	"odonto/internal/domain/appointment"
	"odonto/internal/infra/criteria"
	"odonto/internal/infra/data/pgclient"
	"odonto/internal/infra/fails"
	"odonto/internal/infra/utils"
)

type appointmentManager interface {
	Create(e *appointment.Entity) error
	Update(e *appointment.Entity) error
	List(id appointment.Identity, text string, page, limit int) ([]appointment.EntityProxy, int, error)
	Find(e *appointment.Identity) (*appointment.Entity, error)
	Remove(e *appointment.Identity) error
}

type appointmentManagerImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (s *appointmentManagerImpl) Create(e *appointment.Entity) error {
	// new appointment repository
	rep := pgclient.NewAppointmentRepository(s.ctx, s.tx)
	e.ID.Aid = utils.GenerateUUID()
	e.Status = "pending"
	// store appointment
	err := rep.Store(e)
	if err != nil {
		return fails.ProcessStack(103, ErrOnCreateEntity, err)
	}
	// success
	return nil
}

func (s *appointmentManagerImpl) Find(identity *appointment.Identity) (*appointment.Entity, error) {
	// new appointment repository
	rep := pgclient.NewAppointmentRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.And("aid", criteria.OPT_EQUAL, identity.Aid)
	// result
	ent, err := rep.Find(builder)
	if err != nil {
		return nil, fails.ProcessStack(108, ErrOnFindEntity, err)
	}
	return ent, nil
}

func (s *appointmentManagerImpl) Update(e *appointment.Entity) error {
	// new appointment repository
	rep := pgclient.NewAppointmentRepository(s.ctx, s.tx)
	// update appointment
	err := rep.Update(e)
	if err != nil {
		return fails.ProcessStack(111, ErrOnUpdateEntity, err)
	}
	return nil
}

func (s *appointmentManagerImpl) Remove(identity *appointment.Identity) error {
	// new appointment repository
	rep := pgclient.NewAppointmentRepository(s.ctx, s.tx)
	// delete appointment
	err := rep.Remove(identity)
	if err != nil {
		return fails.ProcessStack(113, ErrOnRemoveEntity, err)
	}
	return nil
}

func (s *appointmentManagerImpl) List(id appointment.Identity, text string, page, limit int) ([]appointment.EntityProxy, int, error) {
	// new appointment repository
	rep := pgclient.NewAppointmentRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.SetPage(page).SetLimit(limit)
	// builder.NewCriteria(&criteria.Condition{
	// 	Key:      "deleted_at",
	// 	Value:    "null",
	// 	Operator: criteria.OPTypeIs,
	// })
	// list appointment
	entities, count, err := rep.List(builder)
	if err != nil {
		return nil, 0, fails.ProcessStack(116, ErrOnListEntities, err)
	}
	return entities, count, nil
}

func NewAppointmentManager(ctx context.Context, tx *sql.Tx) appointmentManager {
	return &appointmentManagerImpl{
		ctx: ctx,
		tx:  tx,
	}
}
