package appl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"odonto/internal/domain/person"
	"odonto/internal/domain/shared"
	"odonto/internal/infra/criteria"
	"odonto/internal/infra/data/pgclient"
	"odonto/internal/infra/fails"
	"odonto/internal/infra/utils"
	"strings"
	"time"
)

type personManager interface {
	Create(name, cpf, email, birthdate string) (*person.Entity, error)
	CreateAll(e []person.Entity) error
	Update(id, name, email string) error
	List(id person.Identity, text string, page, limit int) ([]person.Entity, error)
	Find(e *person.Identity) (*person.Entity, error)
	Remove(e *person.Identity) error
	FindByCpf(cpf string) (*person.Entity, error)
}

type personManagerImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (s *personManagerImpl) validate(name, cpf, email, birthdate string) error {
	errs := []string{}
	err := DefaultSizeValidation(name, "Nome", 1, 255)
	if err != nil {
		errs = append(errs, err.Error())
	}
	//
	_, err = time.Parse(shared.DefaultTimeLayout, birthdate)
	if err != nil {
		errs = append(errs, "Data de nascimento inválida, formato esperado: DD/MM/YYYY")
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ","))
	}
	return nil
}

func (s *personManagerImpl) validateUpdate(name, email string) error {
	errs := []string{}
	err := DefaultSizeValidation(name, "Nome", 1, 255)
	if err != nil {
		errs = append(errs, err.Error())
	}
	//
	if len(email) > 0 {
		err = utils.IsValidMail(email, "Email inválido")
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ","))
	}
	return nil
}

func (s *personManagerImpl) FindByCpf(cpf string) (*person.Entity, error) {
	rep := pgclient.NewPersonRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.And("cpf", criteria.OPT_EQUAL, cpf)
	// result
	ent, err := rep.Find(builder)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (s *personManagerImpl) Create(name, cpf, email, birthdate string) (*person.Entity, error) {
	err := s.validate(name, cpf, email, birthdate)
	if err != nil {
		return nil, fails.ProcessStack(100, err.Error(), nil)
	}
	fmt.Println(">>>", birthdate)
	// new person repository
	rep := pgclient.NewPersonRepository(s.ctx, s.tx)
	bdate, _ := time.Parse(shared.DefaultTimeLayout, birthdate)
	e := &person.Entity{
		Name:      name,
		Cpf:       cpf,
		Email:     email,
		Birthdate: bdate,
		ID: person.Identity{
			Pid: utils.GenerateUUID(),
		},
	}
	// store person
	err = rep.Store(e)
	if err != nil {
		return nil, fails.ProcessStack(103, ErrOnCreateEntity, err)
	}
	// success
	return e, nil
}

func (s *personManagerImpl) CreateAll(entities []person.Entity) error {
	// new person repository
	rep := pgclient.NewPersonRepository(s.ctx, s.tx)
	// store person
	err := rep.StoreAll(entities)
	if err != nil {
		return fails.ProcessStack(105, ErrOnCreateEntity, err)
	}
	// success
	return nil
}

func (s *personManagerImpl) Find(identity *person.Identity) (*person.Entity, error) {
	// new person repository
	rep := pgclient.NewPersonRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.And("pid", criteria.OPT_EQUAL, identity.Pid)
	// result
	ent, err := rep.Find(builder)
	if err != nil {
		return nil, err
	}
	return ent, nil
}

func (s *personManagerImpl) Update(id, name, email string) error {
	// new person repository
	rep := pgclient.NewPersonRepository(s.ctx, s.tx)
	// validate
	err := s.validateUpdate(name, email)
	if err != nil {
		return fails.ProcessStack(105, err.Error(), nil)
	}
	pe, err := s.Find(&person.Identity{
		Pid: id,
	})
	if err != nil {
		return err
	}
	pe.Name = name
	pe.Email = email
	// update person
	err = rep.Update(pe)
	if err != nil {
		return fails.ProcessStack(111, ErrOnUpdateEntity, err)
	}
	return nil
}

func (s *personManagerImpl) Remove(identity *person.Identity) error {
	// new person repository
	rep := pgclient.NewPersonRepository(s.ctx, s.tx)
	// delete person
	err := rep.Remove(identity)
	if err != nil {
		return fails.ProcessStack(113, ErrOnRemoveEntity, err)
	}
	return nil
}

func (s *personManagerImpl) List(id person.Identity, text string, page, limit int) ([]person.Entity, error) {
	// new person repository
	rep := pgclient.NewPersonRepository(s.ctx, s.tx)
	// create criteria
	builder := criteria.NewCriteriaBuilder()
	builder.SetPage(page).SetLimit(limit)
	// list person
	entities, err := rep.List(builder)
	if err != nil {
		return nil, fails.ProcessStack(116, ErrOnListEntities, err)
	}
	return entities, nil
}

func NewPersonManager(ctx context.Context, tx *sql.Tx) personManager {
	return &personManagerImpl{
		ctx: ctx,
		tx:  tx,
	}
}
