package errors

type ErrorCode int

const (
	// 권한 에러 코드
	PasswordMismatchErrorCode ErrorCode = 100
	UserNotFoundErrorCode     ErrorCode = 101
	EmailAlreadyExistsErrorCode    ErrorCode = 102
	UserAlreadyExistsErrorCode   ErrorCode = 103
	InvalidAuthHeaderErrorCode  ErrorCode = 104
	TokenMissingErrorCode          ErrorCode = 105
	UnAuthorizedErrorCode ErrorCode = 106
	RoleUpgradeRequestNotFoundErrorCode ErrorCode = 107
	ForbiddenErrorCode ErrorCode = 108
	BadRoleUpgradeRequestErrorCode ErrorCode = 109

	// 비즈니스 에러 코드
	// 1. Email
	EmailVerificationErrorCode ErrorCode = 1001
	EmailVerificationNotFoundErrorCode ErrorCode = 1002

	// Internal Server Error
	InternalServerErrorCode ErrorCode = 5000
	ShouldBindJsonErrorCode  ErrorCode = 5001
	ParseTemplateErrorCode ErrorCode = 5003
	ApplyTemplateErrorCode ErrorCode = 5004

	// DB Error
	DBErrorCode ErrorCode = 9999
)

var (

	// 권한 에러 코드
	ErrInvalidAuthHeader = NewAppError(400, InvalidAuthHeaderErrorCode, "유효하지 않은 Authorization 헤더입니다.")
	ErrTokenMissing = NewAppError(401, TokenMissingErrorCode, "인증 토큰이 누락되었습니다.")
	ErrPasswordMismatch   = NewAppError(401, PasswordMismatchErrorCode, "비밀번호가 일치하지 않습니다.")
	ErrUserNotFound       = NewAppError(404, UserNotFoundErrorCode, "사용자를 찾을 수 없습니다.")
	ErrUserAlreadyExists   = NewAppError(400, UserAlreadyExistsErrorCode, "이미 존재하는 사용자입니다.")
	ErrEmailAlreadyExists = NewAppError(400, EmailAlreadyExistsErrorCode, "이미 존재하는 이메일입니다.")
	ErrUnauthorized   = NewAppError(401, UnAuthorizedErrorCode, "인증되지 않은 사용자입니다.")
	ErrForbidden = NewAppError(403, ForbiddenErrorCode, "권한이 없습니다.")

	// Email 관련
	ErrEmailVerification = NewAppError(400, EmailVerificationErrorCode, "이메일 인증에 실패했습니다.")
	ErrEmailVerificationAlreadyUsed = NewAppError(400, EmailVerificationErrorCode, "이미 사용된 이메일 인증입니다.")
	ErrEmailVerificationExpired = NewAppError(400, EmailVerificationErrorCode, "이메일 인증 코드가 만료되었습니다.")
	ErrEmailVerificationNotFound = NewAppError(404, EmailVerificationNotFoundErrorCode, "이메일 인증 정보를 찾을 수 없습니다.")
	ErrInvalidResetPasswordRequest = NewAppError(400, EmailVerificationErrorCode, "이메일 인증 정보가 일치하지 않습니다.")

	// System 관련
	ErrShouldBindJson = NewAppError(400, ShouldBindJsonErrorCode, "잘못된 JSON 형식입니다.")

	ErrInternalServer = NewAppError(500, InternalServerErrorCode, "서버 내부 오류가 발생했습니다.")
	ErrDatabase       = NewAppError(500, DBErrorCode, "데이터베이스 처리 중 오류가 발생했습니다.")

	ErrParseTemplate = NewAppError(500, ParseTemplateErrorCode, "템플릿 파싱 중 오류가 발생했습니다.")
	ErrApplyTemplate = NewAppError(500, ApplyTemplateErrorCode, "템플릿 적용 중 오류가 발생했습니다.")

	
	ErrBadRequest = NewAppError(400, ShouldBindJsonErrorCode, "잘못된 요청입니다.")
	ErrInvalidUserID = NewAppError(400, UserNotFoundErrorCode, "올바른 user_id 형태가 아닙니다.")
)