
	-- Table tb_patient
	-- Created in 13 Aug 22 04:52 -0300
	CREATE TABLE IF NOT EXISTS tb_patient (
		person_pid text not null ,
		created_at timestamp not null default CURRENT_TIMESTAMP,
		updated_at timestamp,
		deleted_at timestamp,
		constraint pk_tb_patient primary key(person_pid)
	);
