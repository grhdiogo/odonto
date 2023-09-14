
	-- Table tb_person
	-- Created in 13 Aug 22 04:52 -0300
	CREATE TABLE IF NOT EXISTS tb_person (
		pid text not null ,
		name text not null ,
		cpf text not null ,
		email text not null ,
		birthdate timestamp not null ,
		created_at timestamp not null default CURRENT_TIMESTAMP,
		updated_at timestamp,
		deleted_at timestamp,
		constraint pk_tb_person primary key(pid)
	);
