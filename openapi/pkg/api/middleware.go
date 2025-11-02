package api

import (
	"context"
	"fmt"
	"net/http"

	jwtn "github.com/orenvadi/kuga-lms/internal/lib/jwt"
)

// StrictJWTMiddleware extracts the Bearer token and adds it to context.
func StrictJWTMiddlewareWithSecretKey(secretKey string) StrictMiddlewareFunc {
	return func(next StrictHandlerFunc, operationID string) StrictHandlerFunc {
		if operationID == "PostStudentLogin" || operationID == "PostTeacherLogin" {
			return next
		}
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (any, error) {
			claims, err := jwtn.ValidateToken(r, secretKey)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				switch operationID {

				case "PostStudentScan":
					return PostStudentScan401JSONResponse{Error: &invalidJwt}, fmt.Errorf("missing or invalid token err: %w", err)
				case "GetStudentSchedule":
					return GetStudentSchedule401JSONResponse{Error: &invalidJwt}, fmt.Errorf("missing or invalid token err: %w", err)
				case "PostTeacherQrStream":
					return PostTeacherQrStream401JSONResponse{Error: &invalidJwt}, fmt.Errorf("missing or invalid token err: %w", err)
				case "GetTeacherSchedule":
					return GetTeacherSchedule401JSONResponse{Error: &invalidJwt}, fmt.Errorf("missing or invalid token err: %w", err)
				}
			}

			userID := claims["uid"].(int64)

			ctx = context.WithValue(ctx, "uid", userID)

			// Proceed to next handler
			return next(ctx, w, r, request)
		}
	}
}
