package appl

import (
	"context"
	"database/sql"
	"odonto/internal/domain/patient"
	"odonto/internal/infra/criteria"
	"odonto/internal/infra/data/pgclient"
	"odonto/internal/infra/fails"
)

type patientManager interface {
	Create(name, cpf, email, birthdate string) error
	Update(id, name, email string) error
	List(text string, page, limit int) ([]patient.Entity, int, error)
	Find(e *patient.Identity) (*patient.Entity, error)
	Remove(e *patient.Identity) error
}

type patientManagerImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (s *patientManagerImpl) Create(name, cpf, email, birthdate string) error {
	// find person by cpf
	pm := NewPersonManager(s.ctx, s.tx)
	// create person
	pe, err := pm.Create(name, cpf, email, birthdate)
	if err != nil {
		return fails.ProcessStack(104, "Falha ao criar pessoa", err)
	}
	// new patient repository
	rep := pgclient.NewPatientRepository(s.ctx, s.tx)
	e := &patient.Entity{
		ID: patient.Identity{
			PersonPid: pe.ID.Pid,
		},
	}
	// verify if person is already a patient
	doc, err := s.Find(&e.ID)
	if err != nil && err != pgclient.ErrFindPatient {
		return fails.ProcessStack(99, "Falha ao procurar paciente", err)
	}
	if doc != nil {
		fails.ProcessStack(100, "Cpf já utilizado para este paciente", nil)
	}
	// store patient
	err = rep.Store(e)
	if err != nil {
		return fails.ProcessStack(555, ErrOnCreateEntity, err)
	}
	// success
	return nil
}

func (s *patientManagerImpl) Find(identity *patient.Identity) (*patient.Entity, error) {
	// new patient repository
	rep := pgclient.NewPatientRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.And("patient.person_pid", criteria.OPT_EQUAL, identity.PersonPid)
	// result
	ent, err := rep.Find(builder)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (s *patientManagerImpl) Update(id, name, email string) error {
	// verify if is patient
	_, err := s.Find(&patient.Identity{
		PersonPid: id,
	})
	if err != nil {
		return err
	}
	pm := NewPersonManager(s.ctx, s.tx)
	// update patient
	err = pm.Update(id, name, email)
	if err != nil {
		return fails.ProcessStack(111, ErrOnUpdateEntity, err)
	}
	return nil
}

func (s *patientManagerImpl) Remove(identity *patient.Identity) error {
	// new patient repository
	rep := pgclient.NewPatientRepository(s.ctx, s.tx)
	// delete patient
	err := rep.Remove(identity)
	if err != nil {
		return fails.ProcessStack(113, ErrOnRemoveEntity, err)
	}
	return nil
}

func (s *patientManagerImpl) List(text string, page, limit int) ([]patient.Entity, int, error) {
	// new patient repository
	rep := pgclient.NewPatientRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.SetPage(page).SetLimit(limit)
	// list patient
	entities, count, err := rep.List(builder)
	if err != nil {
		return nil, 0, fails.ProcessStack(116, ErrOnListEntities, err)
	}
	return entities, count, nil
}

func NewPatientManager(ctx context.Context, tx *sql.Tx) patientManager {
	return &patientManagerImpl{
		ctx: ctx,
		tx:  tx,
	}
}
