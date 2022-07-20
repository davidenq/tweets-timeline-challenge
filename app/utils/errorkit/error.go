package errorkit

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/aws/smithy-go"
)

type ExternalError struct {
	Code int
	Err  interface{}
}

type AppError struct {
	Code int
	Err  interface{}
	*ExternalError
}

func (ke *AppError) Error() string {
	switch dataType := ke.Err.(type) {
	case error:
		return dataType.Error()
	default:
		if ke.Err != nil {
			return ke.Err.(string)
		}
		return ""
	}
}

func NewError(code int, err error) error {
	return &AppError{
		Code: code,
		Err:  err,
	}
}

func NewServerError(message string) error {
	err := errors.New(message)
	return &AppError{
		Code: http.StatusInternalServerError,
		Err:  err,
	}
}

func NewClientError(message string) error {
	err := errors.New(message)
	return &AppError{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}

func NewNetError(response *http.Response, err error) error {
	return statusBadGateway(response.StatusCode, err)
}

func NewDBError(err error) error {
	var apiErr smithy.APIError
	var code int
	if errors.As(err, &apiErr) {
		errCode, _ := strconv.Atoi(apiErr.ErrorCode())
		code = errCode
	}
	return statusBadGateway(code, err)
}

func statusBadGateway(code int, err error) error {
	return &AppError{
		Code: http.StatusBadGateway,
		Err:  http.StatusText(http.StatusBadGateway),
		ExternalError: &ExternalError{
			Code: code,
			Err:  err,
		},
	}
}
