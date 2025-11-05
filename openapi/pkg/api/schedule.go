package api

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s Server) GetStudentSchedule(ctx context.Context, request GetStudentScheduleRequestObject) (GetStudentScheduleResponseObject, error) {
	op := "server.GetStudentSchedule"
	userId := ctx.Value("uid").(string)

	dbScheduleRows, err := s.db.Db.GetStudentSchedule(ctx, userId)
	if err != nil {
		log.Printf("schedule not found, err: %v\n", err)
		return GetStudentSchedule404JSONResponse{Error: &scheduleNotFound}, fmt.Errorf("%s: %w", op, err)
	}

	responseScheduleRows := convertDbStudentScheduleRowsToResponseScheduleRows(dbScheduleRows)

	response := append(GetStudentSchedule200JSONResponse{}, responseScheduleRows...)

	return response, nil
}

func (s Server) GetTeacherSchedule(ctx context.Context, request GetTeacherScheduleRequestObject) (GetTeacherScheduleResponseObject, error) {
	op := "server.GetTeacherSchedule"
	userId := ctx.Value("uid").(string)

	dbScheduleRows, err := s.db.Db.GetTeacherSchedule(ctx, pgtype.Text{String: userId, Valid: true})
	if err != nil {
		log.Printf("schedule not found, err: %v\n", err)
		return GetTeacherSchedule404JSONResponse{Error: &scheduleNotFound}, fmt.Errorf("%s: %w", op, err)
	}

	responseScheduleRows := convertDbTeacherScheduleRowsToResponseScheduleRows(dbScheduleRows)

	response := append(GetTeacherSchedule200JSONResponse{}, responseScheduleRows...)

	return response, nil
}
