
	-- Table tb_appointment
	-- Created in 13 Aug 22 04:52 -0300
	CREATE TABLE IF NOT EXISTS tb_appointment (
		aid text not null ,
		status text not null ,
		observation text not null,
		doctor_did text not null ,
		patient_pid text not null ,
		created_at timestamp not null default CURRENT_TIMESTAMP,
		updated_at timestamp,
		deleted_at timestamp,
		constraint pk_tb_appointment primary key(aid)
	);
