CREATE TABLE users
(
    id serial not null unique,
    email varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE news
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255) not null,
    date timestamp not null,
    user_id int references users (id) on delete cascade not null
);