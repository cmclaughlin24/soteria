package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func SendJsonResponse(w http.ResponseWriter, statusCode int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(out)

	return nil
}

func ReadJsonPayload[T any](r *http.Request, data *T) error {
	decorder := json.NewDecoder(r.Body)

	if err := decorder.Decode(&data); err != nil {
		return err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}
