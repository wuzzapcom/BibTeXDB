CREATE DATABASE bibtex_literature;

DROP SCHEMA schema CASCADE;

CREATE SCHEMA schema;

CREATE TABLE schema.textbook(
  id SERIAL PRIMARY KEY,
  ident VARCHAR(100) NOT NULL,
  title VARCHAR(100) NOT NULL,
  author VARCHAR(100) NOT NULL,
  publisher VARCHAR(100) NOT NULL,
  year INTEGER NOT NULL,
  isbn VARCHAR(20) NOT NULL,
  url VARCHAR(100) NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  timestamp INTEGER NOT NULL
);

CREATE TABLE schema.department(
  department_id SERIAL PRIMARY KEY,
  title VARCHAR(40) UNIQUE NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  timestamp INTEGER NOT NULL
);

CREATE TABLE schema.lecturer(
  lecturer_id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  date_of_birth date NOT NULL,
  department_id SERIAL REFERENCES schema.department,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  timestamp INTEGER NOT NULL,
  UNIQUE (name, date_of_birth)
);

CREATE TABLE schema.course(
  id SERIAL PRIMARY KEY,
  title VARCHAR(40) NOT NULL,
  lecturer_id SERIAL REFERENCES schema.lecturer,
  department_id SERIAL REFERENCES schema.department,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  timestamp INTEGER NOT NULL
);

CREATE TABLE schema.literature_lists(
  id SERIAL PRIMARY KEY,
  year INTEGER,
  course SERIAL REFERENCES schema.course,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  timestamp INTEGER NOT NULL
);

CREATE TABLE schema.literature(
  textbook_id INTEGER REFERENCES schema.textbook,
  literature_list_id SERIAL REFERENCES schema.literature_lists,
  PRIMARY KEY (textbook_id, literature_list_id),
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  timestamp INTEGER NOT NULL
);
