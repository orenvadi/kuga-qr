-- +goose Up
-- Insert test data

-- Groups
INSERT INTO the_group (id) VALUES
('G101'),
('G102');

-- Users (students and teachers share this table)
INSERT INTO the_user (id, full_name, password_hash) VALUES
-- Students
('S1001', 'Alice Johnson', '$2a$10$4icUGNUMcyH5UfOYw2lkLu5d0wYDW65vPVTerzCvMQtdNuC2Vu1Zy'),
('S1002', 'Bob Smith', '$2a$10$4icUGNUMcyH5UfOYw2lkLu5d0wYDW65vPVTerzCvMQtdNuC2Vu1Zy'),
('S1003', 'Carol Davis', '$2a$10$4icUGNUMcyH5UfOYw2lkLu5d0wYDW65vPVTerzCvMQtdNuC2Vu1Zy'),
-- Teachers
('T2001', 'Dr. Emily White', '$2a$10$4icUGNUMcyH5UfOYw2lkLu5d0wYDW65vPVTerzCvMQtdNuC2Vu1Zy'),
('T2002', 'Prof. Michael Brown', '$2a$10$4icUGNUMcyH5UfOYw2lkLu5d0wYDW65vPVTerzCvMQtdNuC2Vu1Zy');

-- Admin (optional, for completeness)
INSERT INTO admin (id, password_hash) VALUES
(1, '$2a$10$adminhash');

-- Students
INSERT INTO student (id, semester, group_id) VALUES
('S1001', 3, 'G101'),
('S1002', 3, 'G101'),
('S1003', 5, 'G102');

-- Teachers
INSERT INTO teacher (id) VALUES
('T2001'),
('T2002');

-- Subjects
INSERT INTO subject (id, teacher_id) VALUES
('MATH101', 'T2001'),
('PHYS201', 'T2002');

-- Rooms
INSERT INTO room (id, room_code) VALUES
('R1', 'A101'),
('R2', 'B205');

-- Courses
INSERT INTO course (id, group_id, subject_id, teacher_id) VALUES
(1, 'G101', 'MATH101', 'T2001'),
(2, 'G101', 'PHYS201', 'T2002'),
(3, 'G102', 'MATH101', 'T2001');

-- Schedules
INSERT INTO schedule (id, subject_id, group_id, teacher_id, room_id, the_time) VALUES
(101, 'MATH101', 'G101', 'T2001', 'R1', '2025-11-03 09:00:00'),
(102, 'PHYS201', 'G101', 'T2002', 'R2', '2025-11-03 11:00:00'),
(103, 'MATH101', 'G102', 'T2001', 'R1', '2025-11-04 10:00:00');

-- QR Codes
INSERT INTO qrcodes (id, code, schedule_id) VALUES
(1, 'QR-MATH101-G101-20251103', 101),
(2, 'QR-PHYS201-G101-20251103', 102),
(3, 'QR-MATH101-G102-20251104', 103);

-- Marks and Absence
INSERT INTO marks_and_absence (schedule_id, student_id, mark, status) VALUES
(101, 'S1001', 85, 'present'),
(101, 'S1002', NULL, 'absent'),
(102, 'S1001', 90, 'present'),
(102, 'S1002', 78, 'present'),
(103, 'S1003', 88, 'present');

-- Colloquium
INSERT INTO colloqium (id, the_year, semester, course_id, colloq_1_score, colloq_2_score, colloq_3_score, exam_score, additional_score, total_score) VALUES
(1, 2025, 3, 1, 25, 24, 23, 28, 5, 105),
(2, 2025, 3, 2, 22, 20, 21, 27, 3, 93);

-- Course Work
INSERT INTO course_work (id, the_year, semester, course_id, score) VALUES
(1, 2025, 3, 1, 92),
(2, 2025, 3, 2, 88);

-- +goose Down
-- Remove test data in reverse dependency order

DELETE FROM course_work;
DELETE FROM colloqium;
DELETE FROM marks_and_absence;
DELETE FROM qrcodes;
DELETE FROM schedule;
DELETE FROM course;
DELETE FROM subject;
DELETE FROM room;
DELETE FROM teacher;
DELETE FROM student;
DELETE FROM admin;
DELETE FROM the_user;
DELETE FROM the_group;
