package api

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	sqlc "github.com/orenvadi/kuga-lms/storage/sql/gen"
)

func (s Server) PostStudentScan(ctx context.Context, request PostStudentScanRequestObject) (PostStudentScanResponseObject, error) {
	op := "server.PostStudentScan"

	studentID, ok := ctx.Value("uid").(string)
	if !ok {
		log.Println("missing student ID in context")
		return PostStudentScan401JSONResponse{Error: &invalidJwt}, fmt.Errorf("%s: missing student ID", op)
	}

	scheduleID := request.Body.ScheduleId
	scannedCode := request.Body.QrData

	// Fetch the expected QR code for this schedule
	qr, err := s.db.Db.GetQrcodeByScheduleID(ctx, pgtype.Int4{Int32: int32(scheduleID), Valid: true})
	if err != nil {
		log.Printf("QR code not found for schedule %d: %v\n", scheduleID, err)
		return PostStudentScan500JSONResponse{Error: &serverError}, fmt.Errorf("%s: invalid schedule", op)
	}

	if qr.Code.String != scannedCode {
		log.Printf("QR code mismatch for schedule %d", scheduleID)
		return PostStudentScan500JSONResponse{Error: &serverError}, fmt.Errorf("%s: invalid QR code", op)
	}

	// Mark attendance: insert or update marks_and_absence
	err = s.db.Db.UpsertAttendance(ctx, sqlc.UpsertAttendanceParams{
		ScheduleID: int32(scheduleID),
		StudentID:  studentID,
		Status:     pgtype.Text{String: "present", Valid: true},
	})
	if err != nil {
		log.Printf("failed to record attendance: %v\n", err)
		return PostStudentScan500JSONResponse{Error: &serverError}, fmt.Errorf("%s: db error", op)
	}

	message := "Attendance recorded successfully"
	return PostStudentScan200JSONResponse{Message: &message}, nil
}

func (s Server) PostTeacherQrStream(ctx context.Context, request PostTeacherQrStreamRequestObject) (PostTeacherQrStreamResponseObject, error) {
	op := "server.PostTeacherQrStream"

	teacherID, ok := ctx.Value("uid").(string)
	if !ok {
		log.Println("missing teacher ID in context")
		return PostTeacherQrStream401JSONResponse{Error: &invalidJwt}, fmt.Errorf("%s: missing teacher ID", op)
	}

	scheduleID := request.Body.ScheduleId

	// Verify this schedule belongs to the teacher
	sched, err := s.db.Db.GetSchedule(ctx, int32(scheduleID))
	if err != nil {
		log.Printf("schedule not found: %v\n", err)
		return PostTeacherQrStream500JSONResponse{Error: &serverError}, fmt.Errorf("%s: schedule not found", op)
	}

	if sched.TeacherID.String != teacherID {
		log.Printf("teacher %s does not own schedule %d\n", teacherID, scheduleID)
		return PostTeacherQrStream401JSONResponse{Error: &invalidJwt}, fmt.Errorf("%s: forbidden", op)
	}

	// Get or create QR code for this schedule
	var qrCode string
	qr, err := s.db.Db.GetQrcodeByScheduleID(ctx, pgtype.Int4{Int32: int32(scheduleID), Valid: true})
	if err != nil {
		// Generate a simple deterministic QR payload
		qrCode = fmt.Sprintf("QR-%s-%s-%d", sched.SubjectID.String, sched.GroupID.String, scheduleID)
		// Optionally insert into DB (if idempotent)
		_, err := s.db.Db.CreateQrcode(ctx, sqlc.CreateQrcodeParams{
			Code:       pgtype.Text{String: qrCode, Valid: true},
			ScheduleID: pgtype.Int4{Int32: int32(scheduleID), Valid: true},
		})
		if err != nil {
			log.Printf("failed to create QR code: %v\n", err)
		}
	} else {
		qrCode = qr.Code.String
	}

	// Simulate an SSE stream: "data: <qr_code>\n\n"
	event := fmt.Sprintf("data: %s\n\n", qrCode)
	reader := bytes.NewReader([]byte(event))

	return PostTeacherQrStream200TexteventStreamResponse{
		Body:          reader,
		ContentLength: int64(reader.Len()),
	}, nil
}
