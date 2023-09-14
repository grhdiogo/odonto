
	-- Table tb_appointment_procedure
	-- Created in 13 Aug 22 04:52 -0300
	CREATE TABLE IF NOT EXISTS tb_appointment_procedure (
		appointment_id text not null ,
		procedure_pid text not null,
		procedure_name text not null,
		procedure_value real not null ,
		tooth int not null,
		created_at timestamp not null default CURRENT_TIMESTAMP,
		updated_at timestamp,
		deleted_at timestamp
	);
