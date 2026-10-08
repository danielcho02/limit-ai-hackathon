package errors

import (
	"fmt"
)

type AppError struct {
	Code     ErrorCode `json:"code"`    // 에러 코드
	Message  string    `json:"message"` // 사용자에게 노출할 메시지
	HTTPCode int       `json:"-"`       // HTTP 상태 코드
	RawError error     `json:"-"`       // 에러 원본

}

// 기존 Error 인터페이스 구현
func (e AppError) Error() string {
	if e.RawError != nil {
		return fmt.Sprintf("[%d] %s (Raw: %v)", e.Code, e.Message, e.RawError)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.RawError
}

func (e *AppError) Wrap(err error) *AppError {
	newErr := *e
	newErr.RawError = err
	return &newErr
}

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func NewAppError(
	httpCode int,
	code ErrorCode,
	message string,
) *AppError {
	return &AppError{
		Code:     code,
		Message:  message,
		HTTPCode: httpCode,
	}
}

// AppError 인지 체크하는 함수
func IsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}