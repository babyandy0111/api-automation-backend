package er

const (
	ErrorParamInvalid     = 400400
	UnauthorizedError     = 400401
	ForbiddenError        = 400403
	ResourceNotFoundError = 400404
	TokenExpiredError     = 401001
	DataDuplicateError    = 400001
	LimitExceededError    = 400002
	UnknownError          = 500000
	DBInsertError         = 500001
	DBUpdateError         = 500002
	DBDeleteError         = 500003
	DBDuplicateKeyError   = 500004
	RedisSerError         = 500005
)
