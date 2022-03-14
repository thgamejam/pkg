// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

func ErrorInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsIncorrectAccount(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INCORRECT_ACCOUNT.String() && e.Code == 500
}

func ErrorIncorrectAccount(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_INCORRECT_ACCOUNT.String(), fmt.Sprintf(format, args...))
}
