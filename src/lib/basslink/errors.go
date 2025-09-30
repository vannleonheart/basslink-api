package basslink

const (
	ErrBadRequest               = "400"
	ErrBadRequestInvalidRequest = "00"
	ErrBadRequestValidation     = "01"

	ErrUnauthorized                     = "401"
	ErrUnauthorizedNeedAuthentication   = "00"
	ErrUnauthorizedAuthenticationFailed = "01"
	ErrUnauthorizedUserDisabled         = "02"
	ErrUnauthorizedUserInvalidType      = "03"
	ErrUnauthorizedUserNotVerified      = "04"
	ErrUnauthorizedTokenInvalid         = "05"
	ErrUnauthorizedInvalidPassword      = "06"

	ErrForbidden                     = "403"
	ErrForbiddenRouteNotPermitted    = "00"
	ErrForbiddenResourceNotPermitted = "01"

	ErrNotFound                 = "404"
	ErrNotFoundRouteNotExist    = "00"
	ErrNotFoundResourceNotExist = "01"

	ErrConflict              = "409"
	ErrConflictResourceExist = "00"
	ErrConflictResourceState = "01"

	ErrInternalHost                = "500"
	ErrInternalHostUnknown         = "00"
	ErrInternalHostDatabase        = "01"
	ErrInternalHostExternalService = "02"
	ErrInternalHostInternalLibrary = "03"
)

type AppError struct {
	Code     string
	Kind     string
	Data     interface{}
	internal string
	message  string
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) GetInternalMessage() string {
	return e.internal
}

func NewAppError(message, code, kind, internal string, data interface{}) *AppError {
	return &AppError{
		Code:     code,
		Kind:     kind,
		Data:     data,
		internal: internal,
		message:  message,
	}
}
