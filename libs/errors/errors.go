package errors

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code    int
	Message string
}

const (
	Unknown       = codes.Unknown         // 2
	Invalid       = codes.InvalidArgument // 3
	NotFound      = codes.NotFound        // 5
	AlreadyExists = codes.AlreadyExists   // 6
	Generic       = codes.Internal        // 13
	Internal      = codes.Internal        // 13
)

var httpCodes = map[codes.Code]int{
	codes.Unknown:          http.StatusConflict,
	codes.Internal:         http.StatusInternalServerError,
	codes.InvalidArgument:  http.StatusBadRequest,
	codes.NotFound:         http.StatusNotFound,
	codes.AlreadyExists:    http.StatusConflict,
	codes.PermissionDenied: http.StatusUnauthorized,
	codes.Unauthenticated:  http.StatusForbidden,
}

func NewError(message string) *Error {
	return &Error{Code: http.StatusInternalServerError, Message: message}
}

func NewGrpcError(c codes.Code, message string) error {
	// i := uint32(code)
	// c := codes.Code(i)
	return status.Error(c, message)
}

func StatusCodeFrom(err error) int {
	if err != nil {
		grpcCode, _ := status.FromError(err)
		httpCode := httpCodes[grpcCode.Code()]
		return httpCode
	}
	return http.StatusConflict
}

func GetCodeFrom(err error) codes.Code {
	if err != nil {
		st, _ := status.FromError(err)
		return st.Code()
	}
	return Unknown
}

func NewErrorFrom(err error) *Error {
	if err != nil {
		grpcCode, _ := status.FromError(err)
		httpCode := httpCodes[grpcCode.Code()]
		return &Error{Code: httpCode, Message: err.Error()}
	}
	return nil
}

func (this *Error) Error() string {
	return fmt.Sprintf("code: %d. message: %s", this.Code, this.Message)
}
