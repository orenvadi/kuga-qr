-- name: GetQrcodeByScheduleID :one
SELECT id, code, schedule_id FROM qrcodes WHERE schedule_id = $1;

-- name: UpsertAttendance :exec
INSERT INTO marks_and_absence (schedule_id, student_id, status)
VALUES ($1, $2, $3)
ON CONFLICT (schedule_id, student_id)
DO UPDATE SET status = EXCLUDED.status;
--
-- name: GetSchedule :one
SELECT id, subject_id, group_id, room_id, teacher_id, the_time FROM schedule WHERE id = $1;

-- name: CreateQrcode :one
INSERT INTO qrcodes (code, schedule_id) VALUES ($1, $2) RETURNING id, code, schedule_id;
