package base

var (
	InvalidParamErr   = NewException(400, "InvalidParam", "An invalid request payload is supplied.")
	ResourceNotFound  = NewException(404, "ResourceNotFound", "The specified resource is not found.")
	InternalError     = NewException(500, "InternalError", "There is an internal error occurred.")
	ResourceInUse     = NewException(409, "ResourceInUse", "The specified resource already exists.")
	RateLimitExceeded = NewException(429, "RateLimitExceeded", "Request is due to rate limit.")
	PermissionDenied  = NewException(403, "PermissionDenied", "You have no permission to do this operation.")
	Unauthorized      = NewException(401, "Unauthorized", "Unauthorized identity.")

	ExceedMaximumKongServiceName = NewException(400, "InvalidParam.ExceedMaximumKongServiceName", "The specified name is too long.")
	InvalidKongServiceId         = NewException(400, "InvalidParam.InvalidKongServiceId", "The specified id is invalid.")

	NotFoundKongService = NewException(404, "ResourceNotFound.NotFoundKongService", "The specified kong service is not found.")
	PasswordLengthErr   = NewException(400, "InvalidParam.PasswordLengthErr", "The length of password must be controlled by 8-20.")
	EmailFormatErr      = NewException(400, "InvalidParam.EmailFormatErr", "The email format is invalid.")

	AuthenticationDenied  = NewException(401, "Unauthorized.AuthenticationDenied", "Authentication denied.")   // expect retry
	AuthenticationExpired = NewException(440, "Unauthorized.AuthenticationExpired", "Authentication expired.") // expect retry
)
