-- +goose Up
-- Create the user table
CREATE TABLE the_user (
    id VARCHAR PRIMARY KEY,
    full_name VARCHAR,
    password_hash VARCHAR
);

-- Create the admin table
CREATE TABLE admin (
    id INT PRIMARY KEY,
    password_hash VARCHAR
);

-- Create the group table
CREATE TABLE the_group (
    id VARCHAR PRIMARY KEY
);

-- Create the student table with foreign keys to user and group
CREATE TABLE student (
    id VARCHAR PRIMARY KEY,
    semester INT,
    group_id VARCHAR,
    FOREIGN KEY (id) REFERENCES the_user(id),
    FOREIGN KEY (group_id) REFERENCES the_group(id)
);

-- Create the teacher table with foreign key to user
CREATE TABLE teacher (
    id VARCHAR PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES the_user(id)
);

-- Create the subject table with foreign key to teacher
CREATE TABLE subject (
    id VARCHAR PRIMARY KEY,
    teacher_id VARCHAR,
    FOREIGN KEY (teacher_id) REFERENCES teacher(id)
);

-- Create the room table
CREATE TABLE room (
    id VARCHAR PRIMARY KEY,
    room_code VARCHAR
);

-- Create the schedule table with foreign keys to subject, group, teacher, and room
CREATE TABLE schedule (
    id INT PRIMARY KEY,
    subject_id VARCHAR,
    group_id VARCHAR,
    teacher_id VARCHAR,
    room_id VARCHAR,
    the_time TIMESTAMP,
    FOREIGN KEY (subject_id) REFERENCES subject(id),
    FOREIGN KEY (group_id) REFERENCES the_group(id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(id),
    FOREIGN KEY (room_id) REFERENCES room(id)
);

-- Create the course table with foreign keys to group, subject, and teacher
CREATE TABLE course (
    id INT PRIMARY KEY,
    group_id VARCHAR,
    subject_id VARCHAR,
    teacher_id VARCHAR,
    FOREIGN KEY (group_id) REFERENCES the_group(id),
    FOREIGN KEY (subject_id) REFERENCES subject(id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(id)
);

-- Create the marks_and_absence table with composite primary key and foreign keys to schedule and student
CREATE TABLE marks_and_absence (
    schedule_id INT,
    student_id VARCHAR,
    mark INT,
    status VARCHAR,
    PRIMARY KEY (schedule_id, student_id),
    FOREIGN KEY (schedule_id) REFERENCES schedule(id),
    FOREIGN KEY (student_id) REFERENCES student(id)
);

-- Create the colloqium table with foreign key to course
CREATE TABLE colloqium (
    id INT PRIMARY KEY,
    the_year INT,
    semester INT,
    course_id INT,
    colloq_1_score INT,
    colloq_2_score INT,
    colloq_3_score INT,
    exam_score INT,
    additional_score INT,
    total_score INT,
    FOREIGN KEY (course_id) REFERENCES course(id)
);

-- Create the course_work table with foreign key to course
CREATE TABLE course_work (
    id INT PRIMARY KEY,
    the_year INT,
    semester INT,
    course_id INT,
    score INT,
    FOREIGN KEY (course_id) REFERENCES course(id)
);

-- +goose Down
DROP TABLE colloqium;
DROP TABLE course_work;
DROP TABLE marks_and_absence;
DROP TABLE schedule;
DROP TABLE course;
DROP TABLE subject;
DROP TABLE student;
DROP TABLE teacher;
DROP TABLE the_group;
DROP TABLE room;
DROP TABLE the_user;
DROP TABLE admin;
