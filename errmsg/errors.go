package errmsg

// Use this msg for pkg/errors WithMessage to wrap err msg.
// Like:
//
const (
	// Http error msg
	HttpBadRequest         = "bad request"
	HttpInvalidRequestBody = "invalid request body"

	// Json error msg
	JsonMarshalFailed   = "json marshal failed"
	JsonUnmarshalFailed = "json unmarshal failed"

	// Form validate error

)
