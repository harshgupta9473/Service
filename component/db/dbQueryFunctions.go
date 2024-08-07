package db


func (s *PostgresStore) CreateUsersTable() error {
	query := `create table if not exists users(
	id integer generated always as identity primary key,
	email varchar(255) unique not null,
	userID varchar(100) unique not null,
	firstname varchar(100),
	lastname varchar(100),
	phoneNumber serial,
	interests varchar(1000),
	verified boolean default false
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateEmalVerificationTable() error {
	query := `create table if not exists emailverification(
	user_id varchar(100),
	token varchar(100) not null,
	expires_at timestamp not null,
	foreign key (user_id) references users(userID)
	)`
	_, err := s.db.Exec(query)
	return err
}