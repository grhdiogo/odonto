	CREATE TABLE IF NOT EXISTS tb_auth_access (
		token text not null ,
		user_uid text not null ,
		expired_at timestamp,
		created_at timestamp not null default CURRENT_TIMESTAMP,
		updated_at timestamp,
		deleted_at timestamp,
		constraint pk_tb_auth_access primary key(token)
	);
