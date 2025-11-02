package api

import (
	sqlc "github.com/orenvadi/kuga-lms/storage/sql/gen"
)

func convertDbTeacherScheduleRowsToResponseScheduleRows(scheduleRows []sqlc.GetTeacherScheduleRow) []Schedule {
	responseScheduleRows := []Schedule{}

	for _, scheduleRow := range scheduleRows {
		int32Id := scheduleRow.ID
		intId := int(int32Id)
		responseScheduleRows = append(responseScheduleRows, Schedule{
			GroupId:   &scheduleRow.GroupID.String,
			Id:        &intId,
			RoomId:    &scheduleRow.GroupID.String,
			SubjectId: &scheduleRow.SubjectID.String,
			TeacherId: &scheduleRow.TeacherID.String,
			Time:      &scheduleRow.TheTime.Time,
		})
	}
	return responseScheduleRows
}

func convertDbStudentScheduleRowsToResponseScheduleRows(scheduleRows []sqlc.GetStudentScheduleRow) []Schedule {
	responseScheduleRows := []Schedule{}

	for _, scheduleRow := range scheduleRows {
		int32Id := scheduleRow.ID
		intId := int(int32Id)
		responseScheduleRows = append(responseScheduleRows, Schedule{
			GroupId:   &scheduleRow.GroupID.String,
			Id:        &intId,
			RoomId:    &scheduleRow.GroupID.String,
			SubjectId: &scheduleRow.SubjectID.String,
			TeacherId: &scheduleRow.TeacherID.String,
			Time:      &scheduleRow.TheTime.Time,
		})
	}
	return responseScheduleRows
}
