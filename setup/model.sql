create database temporary_redis;

create table users (
	user_id serial not null primary key,
	username varchar(64) not null,
	email varchar(128) not null,
	password varchar(60) not null,
	created_at timestamp with time zone default current_timestamp
);