package api

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	jwtn "github.com/orenvadi/kuga-lms/internal/lib/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s Server) PostStudentLogin(ctx context.Context, request PostStudentLoginRequestObject) (PostStudentLoginResponseObject, error) {
	op := "server.PostStudentLogin"
	user, err := s.db.Db.GetUser(ctx, pgtype.Text{String: request.Body.Id, Valid: true})
	if err != nil {
		log.Printf("user not found, err: %v\n", err)
		return PostStudentLogin404JSONResponse{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(request.Body.Password)); err != nil {
		log.Printf("invalid credentials, err: %v\n", err)

		return PostStudentLogin404JSONResponse{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Println("user logged in successfully")

	accessToken, err := jwtn.NewToken(user, s.tokenTTL, s.jwtSecret)
	if err != nil {
		log.Printf("failed to generate token, err: %v\n", err)

		return PostStudentLogin500JSONResponse{}, fmt.Errorf("%s: %w", op, errors.New("unable to generate jwt token"))
	}

	return PostStudentLogin200JSONResponse{Token: &accessToken}, nil
}

func (s Server) PostTeacherLogin(ctx context.Context, request PostTeacherLoginRequestObject) (PostTeacherLoginResponseObject, error) {
	op := "server.PostTeacherLogin"
	user, err := s.db.Db.GetUser(ctx, pgtype.Text{String: request.Body.Id, Valid: true})
	if err != nil {
		log.Printf("user not found, err: %v\n", err)
		return PostTeacherLogin404JSONResponse{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash.String), []byte(request.Body.Password)); err != nil {
		log.Printf("invalid credentials, err: %v\n", err)

		return PostTeacherLogin404JSONResponse{}, fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Println("user logged in successfully")

	accessToken, err := jwtn.NewToken(user, s.tokenTTL, s.jwtSecret)
	if err != nil {
		log.Printf("failed to generate token, err: %v\n", err)

		return PostTeacherLogin500JSONResponse{}, fmt.Errorf("%s: %w", op, errors.New("unable to generate jwt token"))
	}

	return PostTeacherLogin200JSONResponse{Token: &accessToken}, nil
}
