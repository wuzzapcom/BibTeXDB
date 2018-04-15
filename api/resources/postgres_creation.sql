CREATE DATABASE bibtex_literature;

DROP SCHEMA schema CASCADE;

CREATE SCHEMA schema;

CREATE TABLE schema.textbook(
  textbook_id SERIAL PRIMARY KEY,
  textbook_ident VARCHAR(100) NOT NULL UNIQUE,
  textbook_title VARCHAR(100) NOT NULL,
  textbook_author VARCHAR(100) NOT NULL,
  textbook_publisher VARCHAR(100) NOT NULL,
  textbook_year INTEGER NOT NULL,
  textbook_isbn VARCHAR(20) NOT NULL UNIQUE,
  textbook_url VARCHAR(100) NOT NULL,
  textbook_is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  textbook_timestamp INTEGER NOT NULL
);

CREATE TABLE schema.department(
  department_id SERIAL PRIMARY KEY,
  department_title VARCHAR(40) UNIQUE NOT NULL,
  department_is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  department_timestamp INTEGER NOT NULL
);

CREATE TABLE schema.lecturer(
  lecturer_id SERIAL PRIMARY KEY,
  lecturer_name VARCHAR(40) NOT NULL,
  lecturer_date_of_birth date NOT NULL,
  lecturer_department_id SERIAL REFERENCES schema.department,
  lecturer_is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  lecturer_timestamp INTEGER NOT NULL,
  UNIQUE (lecturer_name, lecturer_date_of_birth)
);

CREATE TABLE schema.course(
  course_id SERIAL PRIMARY KEY,
  course_title VARCHAR(40) NOT NULL,
  course_lecturer_id SERIAL REFERENCES schema.lecturer,
  course_department_id SERIAL REFERENCES schema.department,
  course_semester INTEGER NOT NULL,
  course_is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  course_timestamp INTEGER NOT NULL,
  UNIQUE (course_title, course_department_id, course_semester)
);

CREATE TABLE schema.literature_lists(
  literature_list_id SERIAL PRIMARY KEY,
  literature_list_year INTEGER,
  literature_list_course_id SERIAL REFERENCES schema.course,
  literature_list_is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  literature_list_timestamp INTEGER NOT NULL,
  UNIQUE (literature_list_year, literature_list_course_id)
);

CREATE TABLE schema.literature(
  literature_textbook_id INTEGER REFERENCES schema.textbook,
  literature_literature_list_id SERIAL REFERENCES schema.literature_lists,
  PRIMARY KEY (literature_textbook_id, literature_literature_list_id),
  literature_is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  literature_timestamp INTEGER NOT NULL
);

CREATE USER api_user WITH PASSWORD 'pass';
DROP USER api_user;

REVOKE ALL
ON ALL TABLES IN SCHEMA schema
FROM api_user;

GRANT USAGE ON SCHEMA schema TO api_user;
GRANT SELECT ON ALL TABLES IN SCHEMA schema to api_user;

GRANT SELECT, INSERT, UPDATE
ON ALL TABLES IN SCHEMA schema
TO api_user;

INSERT INTO schema.department(department_title, department_timestamp) VALUES ('ИУ9', 1);

INSERT INTO schema.lecturer(lecturer_name, lecturer_date_of_birth, lecturer_department_id, lecturer_timestamp) VALUES
  ('Лапатин', '1997-02-20', 1, 1);

INSERT INTO schema.course(course_title, course_lecturer_id, course_department_id, course_semester, course_timestamp) VALUES
  ('Компиляторы', 1, 1, 3, 0);

INSERT INTO schema.department(department_title, department_timestamp) VALUES ('Прикладная математика и информатика', 1);

INSERT INTO schema.lecturer(lecturer_name, lecturer_date_of_birth, lecturer_department_id, lecturer_timestamp) VALUES
  ('Скоробогатов Сергей Юрьевич', '2013-02-03', 2, 1);

INSERT INTO schema.literature_lists(literature_list_year, literature_list_course_id, literature_list_timestamp) VALUES
  (2017, 1, 11);

SELECT textbook_ident, literature_list_course_id, literature_list_year FROM schema.literature_lists l
  JOIN (schema.literature
    JOIN schema.textbook
      ON literature.literature_textbook_id = textbook.textbook_id) j
    ON j.literature_literature_list_id = l.literature_list_id;

SELECT literature_list_year, course_title, department_title FROM schema.department d
  JOIN (schema.literature_lists l
    JOIN schema.course c
      ON l.literature_list_course_id = c.course_id) j
    ON d.department_id = j.course_department_id;

SELECT course_title, lecturer_name, department_title, course_semester FROM schema.lecturer
  JOIN (schema.course
    JOIN schema.department
      ON course_department_id = department_id) j
    ON lecturer_id = j.course_lecturer_id;

INSERT INTO schema.literature(literature_textbook_id, literature_literature_list_id, literature_timestamp)
  SELECT literature_textbook_id, 2, literature_timestamp FROM schema.literature WHERE literature_literature_list_id = 1;