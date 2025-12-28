package errorhandler

type UserServiceErrorType string

const (
	NotFound     UserServiceErrorType = "not found"
	Invalid      UserServiceErrorType = "invalid"
	Unknown      UserServiceErrorType = "unknown"
	AlreadyExist UserServiceErrorType = "already exist"
	Internal UserServiceErrorType = "internal error"
)

type UserServiceError struct {
	ErrorType     UserServiceErrorType
	ClientMessage string
	Err           error
}

func (c *UserServiceError) Error() string {
	if c.Err == nil {
		return c.ClientMessage
	}

	return c.Err.Error()
}
