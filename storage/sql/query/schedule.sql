-- name: GetStudentSchedule :many
SELECT id, subject_id, group_id, room_id, teacher_id, the_time
  FROM schedule
  WHERE group_id = (
    SELECT group_id
    FROM student
    WHERE student.id = $1
  );
--
-- name: GetTeacherSchedule :many
SELECT id, subject_id, group_id, room_id, teacher_id, the_time
  FROM schedule
  WHERE teacher_id = $1;
