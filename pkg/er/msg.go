package er

var msg = map[int]string{
	ErrorParamInvalid:     "Wrong parameter format or invalid",
	UnauthorizedError:     "Unauthorized",
	ForbiddenError:        "Forbidden error",
	ResourceNotFoundError: "Resource not found",
	TokenExpiredError:     "Token is expired",
	DataDuplicateError:    "Data duplicate error",
	LimitExceededError:    "Limit exceeded error",
	UnknownError:          "Database unknown error",
	DBInsertError:         "Database insertion error",
	DBUpdateError:         "Database update error",
	DBDeleteError:         "Database delete error",
	DBDuplicateKeyError:   "Database data duplicate key error",
	RedisSerError:         "Redis server error",
}
