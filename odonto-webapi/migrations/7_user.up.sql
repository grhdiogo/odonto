	CREATE TABLE IF NOT EXISTS tb_user (
		person_pid text not null ,
		kind text not null ,
		identifier text not null,
		secret text not null,
		created_at timestamp not null default CURRENT_TIMESTAMP,
		updated_at timestamp,
		deleted_at timestamp,
		constraint pk_tb_user primary key(person_pid)
	);
