package schema

var Schema = `
CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    email varchar(50),
    hash varchar(512),
    token_sugar varchar(50)
);`
