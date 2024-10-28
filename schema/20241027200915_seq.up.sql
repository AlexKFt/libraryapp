CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE books
(
    id serial not null unique,
    author varchar(255) not null, 
    annotation varchar(2048) not null,
    bookpoint varchar(255) not null,
    datetake varchar(255) not null,
    datereturn varchar(255) not null
);