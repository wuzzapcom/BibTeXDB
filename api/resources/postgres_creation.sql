CREATE DATABASE bibtex_literature;

DROP SCHEMA schema;

CREATE SCHEMA schema;

CREATE TABLE schema.textbook(
  id SERIAL PRIMARY KEY,
  ident VARCHAR(100) NOT NULL,
  title VARCHAR(100) NOT NULL,
  author VARCHAR(100) NOT NULL,
  publisher VARCHAR(100) NOT NULL,
  year INTEGER NOT NULL,
  isbn VARCHAR(20) NOT NULL,
  url VARCHAR(100) NOT NULL
);

CREATE TABLE schema.department(
  department_id SERIAL PRIMARY KEY,
  title VARCHAR(40) NOT NULL
);

CREATE TABLE schema.lecturer(
  lecturer_id SERIAL PRIMARY KEY,
  name VARCHAR(40) NOT NULL,
  date_of_birth date NOT NULL,
  department_id SERIAL REFERENCES schema.department
);

CREATE TABLE schema.course(
  id SERIAL PRIMARY KEY,
  title VARCHAR(40) NOT NULL,
  lecturer_id SERIAL REFERENCES schema.lecturer,
  department_id SERIAL REFERENCES schema.department
);

CREATE TABLE schema.literature(
  textbook_id INTEGER REFERENCES schema.textbook,
  literature_list_id SERIAL REFERENCES schema.literature_lists,
  PRIMARY KEY (textbook_id, literature_list_id)
);

CREATE TABLE schema.literature_lists(
  id SERIAL PRIMARY KEY,
  year INTEGER
);

INSERT INTO schema.lecturer(name, date_of_birth, department_id) VALUES ('Trishin', '12.02.1997', 1);
INSERT INTO schema.department(title) VALUES ('UI9');
INSERT INTO schema.course(title, lecturer_id, department_id) VALUES ('Compilers', 1, 1);