package appl

import (
	"context"
	"database/sql"
	"odonto/internal/domain/doctor"
	"odonto/internal/infra/criteria"
	"odonto/internal/infra/data/pgclient"
	"odonto/internal/infra/fails"
)

type doctorManager interface {
	Create(name, cpf, email, birthdate string) error
	Update(id, name, email string) error
	List(text string, page, limit int) ([]doctor.Entity, int, error)
	Find(e *doctor.Identity) (*doctor.Entity, error)
	Remove(e *doctor.Identity) error
}

type doctorManagerImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (s *doctorManagerImpl) Create(name, cpf, email, birthdate string) error {
	// find person by cpf
	pm := NewPersonManager(s.ctx, s.tx)
	// create person
	pe, err := pm.Create(name, cpf, email, birthdate)
	if err != nil {
		return fails.ProcessStack(104, "Falha ao criar pessoa", err)
	}
	// new doctor repository
	rep := pgclient.NewDoctorRepository(s.ctx, s.tx)
	e := &doctor.Entity{
		ID: doctor.Identity{
			PersonPid: pe.ID.Pid,
		},
	}
	// verify if person is already a doctor
	doc, err := s.Find(&e.ID)
	if err != nil && err != pgclient.ErrFindDoctor {
		return fails.ProcessStack(99, "Falha ao procurar doutor", err)
	}
	if doc != nil {
		fails.ProcessStack(100, "Cpf já utilizado para este doutor", nil)
	}
	// store doctor
	err = rep.Store(e)
	if err != nil {
		return fails.ProcessStack(555, ErrOnCreateEntity, err)
	}
	// success
	return nil
}

func (s *doctorManagerImpl) Find(identity *doctor.Identity) (*doctor.Entity, error) {
	// new doctor repository
	rep := pgclient.NewDoctorRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.And("doctor.person_pid", criteria.OPT_EQUAL, identity.PersonPid)
	// result
	ent, err := rep.Find(builder)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (s *doctorManagerImpl) Update(id, name, email string) error {
	// verify if is doctor
	_, err := s.Find(&doctor.Identity{
		PersonPid: id,
	})
	if err != nil {
		return err
	}
	pm := NewPersonManager(s.ctx, s.tx)
	// update doctor
	err = pm.Update(id, name, email)
	if err != nil {
		return fails.ProcessStack(111, ErrOnUpdateEntity, err)
	}
	return nil
}

func (s *doctorManagerImpl) Remove(identity *doctor.Identity) error {
	// new doctor repository
	rep := pgclient.NewDoctorRepository(s.ctx, s.tx)
	// delete doctor
	err := rep.Remove(identity)
	if err != nil {
		return fails.ProcessStack(113, ErrOnRemoveEntity, err)
	}
	return nil
}

func (s *doctorManagerImpl) List(text string, page, limit int) ([]doctor.Entity, int, error) {
	// new doctor repository
	rep := pgclient.NewDoctorRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.SetPage(page).SetLimit(limit)
	// list doctor
	entities, count, err := rep.List(builder)
	if err != nil {
		return nil, 0, fails.ProcessStack(116, ErrOnListEntities, err)
	}
	return entities, count, nil
}

func NewDoctorManager(ctx context.Context, tx *sql.Tx) doctorManager {
	return &doctorManagerImpl{
		ctx: ctx,
		tx:  tx,
	}
}
