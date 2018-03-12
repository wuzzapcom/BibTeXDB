CREATE DATABASE bibtex_literature;

DROP SCHEMA schema;

CREATE SCHEMA schema;

CREATE TABLE schema.textbook(
  id SERIAL PRIMARY KEY,
  ident VARCHAR(100) NOT NULL,
  title VARCHAR(100) NOT NULL,
  author VARCHAR(100) NOT NULL,
  publisher VARCHAR(100) NOT NULL,
  year INT NOT NULL,
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
  textbook_id integer  REFERENCES schema.textbook,
  course_id integer REFERENCES schema.course,
  PRIMARY KEY (textbook_id, course_id)
);