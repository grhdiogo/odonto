
	-- Table tb_procedure
	-- Created in 13 Aug 22 04:52 -0300
	CREATE TABLE IF NOT EXISTS tb_procedure(
		pid text not null ,
		name text not null ,
		value real not null ,
		created_at timestamp not null default CURRENT_TIMESTAMP,
		updated_at timestamp,
		deleted_at timestamp,
		constraint pk_tb_procedure primary key(pid)
	);
