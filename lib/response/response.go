package response

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Renos-id/go-starter-template/lib"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func SetLoggerForResponse(logr *logrus.Logger) {
	logger = logr
}

type CommonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Note    string `json:"-"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
	Code    int    `json:"-"`
}

// respondwithJSON write json response format
func (cr CommonResponse) ToJSON(w http.ResponseWriter) {
	response, _ := json.Marshal(cr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(cr.Code)
	w.Write(response)
	return
}

func WriteSuccess(message string, data any) CommonResponse {
	return CommonResponse{
		Status:  true,
		Message: message,
		Data:    data,
		Code:    200,
	}
}

func WriteError(code int, note string, err any) CommonResponse {
	// var errors ValidationErrors
	if code == 0 {
		code = 500
	}
	var body = CommonResponse{
		Status:  false,
		Message: http.StatusText(code),
		Note:    note,
		Code:    code,
	}

	switch e := err.(type) {
	case ValidationErrors:
		body.Error = err
	case error:
		switch e.(error) {
		case sql.ErrNoRows:
			body.Message = "Data does not exists!"
		case io.EOF:
			body.Message = "Failed to read HTTP Request"
		default:
			body.Message = e.Error()
		}
	default:
		body.Error = err
	}

	logger.WithFields(logrus.Fields{
		"env":          os.Getenv("APP_ENV"),
		"err":          err,
		"note":         note,
		"api_response": lib.StructToJSON(body),
	}).Error("Chat Service Error: " + note)

	return body
}
