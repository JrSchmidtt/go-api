CREATE DATABASE IF NOT EXISTS api;
use api;

DROP TABLE IF EXISTS user;

CREATE TABLE user(
    id int auto_increment primary key unique,
    name varchar(200) not null,
    username varchar(200) not null unique,
    email varchar(200) not null unique,
    password varchar(200) not null unique,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;