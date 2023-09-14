package pgclient

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"odonto/internal/infra/criteria"

	"odonto/internal/domain/appointment"
)

var (
	ErrSaveAppointment   = errors.New("error on save appointment")
	ErrFindAppointment   = errors.New("error on find appointment")
	ErrListAppointment   = errors.New("error on list appointment")
	ErrUpdateAppointment = errors.New("error on update appointment")
	ErrRemoveAppointment = errors.New("error on remove appointment")
)

type appointmentRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

// ============================================================================
//
// ============================================================================

func (r *appointmentRepositoryImpl) scanIteratorItem(rows *sql.Rows) (*appointment.Item, error) {
	//
	name := sql.NullString{}
	value := sql.NullFloat64{}
	tooth := sql.NullInt32{}
	id := sql.NullString{}
	// scan columns
	err := rows.Scan(
		&id,
		&name,
		&value,
		&tooth,
	)
	if err != nil {
		return nil, err
	}
	// create entity
	entity := new(appointment.Item)
	if id.Valid {
		entity.ProcedurePID = id.String
	}
	if name.Valid {
		entity.ProcedureName = name.String
	}
	if value.Valid {
		entity.ProcedureValue = value.Float64
	}
	if tooth.Valid {
		entity.Tooth = int(tooth.Int32)
	}
	// result
	return entity, nil
}

func (r *appointmentRepositoryImpl) listItems(id string) ([]appointment.Item, error) {
	// sql text
	sqlText := `
		select 
			procedure_pid,
			procedure_name,
			procedure_value,
			tooth
		from tb_appointment_procedure
	`

	// execute query in database
	rows, err := r.tx.Query(sqlText, id)
	if err != nil {
		return nil, err
	}
	// iterate
	result := make([]appointment.Item, 0)
	for rows.Next() {
		value, err := r.scanIteratorItem(rows)
		if err != nil {
			rows.Close()
			return nil, err
		}
		result = append(result, *value)
	}
	rows.Close()
	// success
	return result, nil
}

func (r *appointmentRepositoryImpl) storeItem(id string, e *appointment.Item) error {
	// create sql
	sqlText := `
		insert into tb_appointment_procedure ( 
			appointment_id,
			procedure_pid,
			procedure_name,
			procedure_value,
			tooth
		)
		values (
			$1 , $2 , $3 , $4, $5
		) `

	result, err := r.tx.Exec(sqlText,
		id,
		e.ProcedurePID,
		e.ProcedureName,
		e.ProcedureValue,
		e.Tooth,
	)
	if err != nil {
		return err
	}
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// check if insert was affected
	if aff != 1 {
		return errors.New("Falha ao salvar item")
	}
	// result
	return nil
}

func (r *appointmentRepositoryImpl) delete(id string) error {
	// create sql
	sqlText := `
		delete from tb_appointment_procedure where appointment_id = $1
		`

	_, err := r.tx.Exec(sqlText,
		id,
	)
	if err != nil {
		return err
	}
	// result
	return nil
}

// ============================================================================
//
// ============================================================================

func (r *appointmentRepositoryImpl) scanIteratorGeneral(rows *sql.Rows, wc, isProxy bool) (*appointment.EntityProxy, int, error) {
	//
	aid := sql.NullString{}
	status := sql.NullString{}
	observation := sql.NullString{}
	doctorDid := sql.NullString{}
	patientPid := sql.NullString{}
	doctorName := sql.NullString{}
	patientName := sql.NullString{}
	createdAt := sql.NullTime{}
	count := sql.NullInt64{}

	vars := []any{
		&aid,
		&status,
		&doctorDid,
		&patientPid,
		&observation,
	}
	if isProxy {
		vars = append(vars, &doctorName, &patientName, &createdAt)
	}
	if wc {
		vars = append(vars, &count)
	}
	// scan columns
	err := rows.Scan(vars...)
	if err != nil {
		return nil, 0, err
	}
	// create entity
	entity := new(appointment.EntityProxy)
	//
	if aid.Valid {
		entity.ID.Aid = aid.String
	}
	if status.Valid {
		entity.Status = status.String
	}
	if doctorDid.Valid {
		entity.DoctorDid = doctorDid.String
	}
	if patientPid.Valid {
		entity.PatientPid = patientPid.String
	}
	if observation.Valid {
		entity.Observation = observation.String
	}
	if doctorName.Valid {
		entity.DoctorName = doctorName.String
	}
	if patientName.Valid {
		entity.PatientName = patientName.String
	}
	if createdAt.Valid {
		entity.CreatedAt = createdAt.Time
	}
	// result
	return entity, int(count.Int64), nil
}

func (r *appointmentRepositoryImpl) scanIterator(rows *sql.Rows) (*appointment.Entity, int, error) {
	ent, c, err := r.scanIteratorGeneral(rows, true, false)
	if err != nil {
		return nil, 0, err
	}
	return &ent.Entity, c, nil
}

func (r *appointmentRepositoryImpl) store(e *appointment.Entity) error {
	// create sql
	sqlText := `
		insert into tb_appointment ( 
			aid, 
			status, 
			doctor_did, 
			patient_pid,
			observation
		)
		values (
			$1 , $2 , $3 , $4, $5
		) `

	result, err := r.tx.Exec(sqlText,
		e.ID.Aid,
		e.Status,
		e.DoctorDid,
		e.PatientPid,
		e.Observation,
	)
	if err != nil {
		return err
	}
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// check if insert was affected
	if aff != 1 {
		return ErrSaveAppointment
	}

	// result
	return nil
}

func (r *appointmentRepositoryImpl) update(e *appointment.Entity) error {
	// sql text
	sqlText := ` 
	update tb_appointment set 
		status = $1,
		doctor_did = $2,
		patient_pid = $3,
		observation = $5,
		updated_at = CURRENT_TIMESTAMP
		where 
			aid = $4
	`

	result, err := r.tx.Exec(sqlText,
		e.Status,
		e.DoctorDid,
		e.PatientPid,
		e.ID.Aid,
		e.Observation,
	)
	if err != nil {
		return err
	}
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// check if insert was affected
	if aff != 1 {
		return ErrUpdateAppointment
	}
	// result
	return nil
}

func (r *appointmentRepositoryImpl) remove(id *appointment.Identity) error {
	// sql text
	sqlText := `
	update tb_appointment set
		deleted_at = CURRENT_TIMESTAMP
	where
		deleted_at is null
		and aid = $1
	`

	result, err := r.tx.Exec(sqlText,
		id.Aid,
	)
	if err != nil {
		return err
	}
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// check if insert was affected
	if aff != 1 {
		return ErrRemoveAppointment
	}
	// result
	return nil
}

func (r *appointmentRepositoryImpl) find(crt criteria.Criteria) (*appointment.Entity, error) {
	// sql text
	sqlText := fmt.Sprintf(`
		select 
			aid,
			status,
			doctor_did,
			patient_pid,
			observation
			from tb_appointment
			where deleted_at is null
			%s
			`, crt.Build(1, false))

	// execute query in database
	rows, err := r.tx.Query(sqlText, crt.Values()...)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		ent, _, err := r.scanIterator(rows)
		if err != nil {
			rows.Close()
			return nil, err
		}
		rows.Close()
		return ent, nil
	}
	//retun
	return nil, ErrFindAppointment
}

func (r *appointmentRepositoryImpl) list(crt criteria.Criteria) ([]appointment.EntityProxy, int, error) {
	sqlText := crt.BuildWithCount(`
		select 
			a.aid, 
			a.status, 
			a.doctor_did, 
			a.patient_pid,
			a.observation,
			c.name,
			e.name,
			a.created_at
		from tb_appointment a
		inner join tb_person c on a.doctor_did = c.pid
		inner join tb_person e on a.patient_pid = e.pid
		where a.deleted_at is null
	`, 1)
	// execute query in database
	rows, err := r.tx.Query(sqlText, crt.Values()...)
	if err != nil {
		return nil, 0, err
	}
	// iterate
	result := make([]appointment.EntityProxy, 0)
	count := 0
	for rows.Next() {
		value, c, err := r.scanIteratorGeneral(rows, true, true)
		if err != nil {
			rows.Close()
			return nil, 0, err
		}
		result = append(result, *value)
		count = c
	}
	rows.Close()
	// success
	return result, count, nil
}

func (r *appointmentRepositoryImpl) Store(e *appointment.Entity) error {
	err := r.store(e)
	if err != nil {
		return err
	}
	//
	for i := 0; i < len(e.Items); i++ {
		err = r.storeItem(e.ID.Aid, &e.Items[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *appointmentRepositoryImpl) Update(e *appointment.Entity) error {
	err := r.update(e)
	if err != nil {
		return err
	}
	// delete
	err = r.delete(e.ID.Aid)
	if err != nil {
		return err
	}
	for i := 0; i < len(e.Items); i++ {
		err = r.storeItem(e.ID.Aid, &e.Items[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *appointmentRepositoryImpl) Remove(id *appointment.Identity) error {
	return r.remove(id)
}

func (r *appointmentRepositoryImpl) Find(crt criteria.Criteria) (*appointment.Entity, error) {
	ent, err := r.find(crt)
	if err != nil {
		return nil, err
	}
	//
	it, err := r.listItems(ent.ID.Aid)
	if err != nil {
		return nil, err
	}
	ent.Items = it

	return ent, nil
}

func (r *appointmentRepositoryImpl) List(crt criteria.Criteria) ([]appointment.EntityProxy, int, error) {
	list, count, err := r.list(crt)
	if err != nil {
		return nil, 0, err
	}
	//
	for i := 0; i < len(list); i++ {
		it, err := r.listItems(list[i].ID.Aid)
		if err != nil {
			return nil, 0, err
		}
		list[i].Items = it
	}
	return list, count, nil
}

func NewAppointmentRepository(ctx context.Context, tx *sql.Tx) appointment.Repository {
	return &appointmentRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}
