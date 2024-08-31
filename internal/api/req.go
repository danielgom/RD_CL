package api

import (
	"RD-Clone-NAPI/internal/api/validator"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var errInvalidBody = errors.New("invalid request body")

func shouldBindIntoAndValidate[T any](r *http.Request, val *T, v *validator.CustomValidator) error {
	err := json.NewDecoder(r.Body).Decode(val)
	if err != nil {
		slog.Warn("failed to parse request body", "error", err)
		return errInvalidBody
	}

	err = v.Struct(val)
	if err != nil {
		return fmt.Errorf("failed to validate request body: %w", err)
	}

	return nil
}
