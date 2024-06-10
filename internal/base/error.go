package base

var (
	InvalidParamErr   = NewException(400, "InvalidParam", "An invalid value is supplied..")
	ResourceNotFound  = NewException(404, "ResourceNotFound", "The specified resource is not found.")
	InternalError     = NewException(500, "InternalError", "There is an internal error occurred.")
	ResourceInUse     = NewException(409, "ResourceInUse", "The specified resource already exists.")
	RateLimitExceeded = NewException(429, "RateLimitExceeded", "Request is due to rate limit.")
	PermissionDenied  = NewException(403, "PermissionDenied", "You have no permission to do this operation.")
	Unauthorized      = NewException(401, "Unauthorized", "Unauthorized identity.")

	ExceedMaximumKongServiceName = NewException(400, "InvalidParam.ExceedMaximumKongServiceName", "The specified name is too long.")
	InvalidKongServiceId         = NewException(400, "InvalidParam.InvalidKongServiceId", "The specified id is invalid.")

	NotFoundKongService = NewException(404, "ResourceNotFound.NotFoundKongService", "The specified kong service is not found.")
)
