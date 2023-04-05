package Error

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidParameter = NewErrorCode(codes.InvalidArgument, "invalid argument")
	ErrServerError      = NewErrorCode(codes.Internal, "server error")
	ErrPermissionDenied = NewErrorCode(codes.PermissionDenied, "permission denied")
)

type ErrorCodeSt struct {
	Code    codes.Code
	Message string
	Err     error
}

func NewErrorCode(code codes.Code, message string) ErrorCodeSt {
	return ErrorCodeSt{
		Code:    code,
		Message: message,
		Err:     errors.New(message),
	}
}

func (s ErrorCodeSt) Error(err error) error {
	str := s.Message
	if err != nil {
		str = fmt.Sprintf("%s:%s", s.Message, err.Error())
	}

	return status.Error(s.Code, str)
}
