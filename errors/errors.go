package errors

type ErrorCode interface {
	//Message A human friendly message reprensenting what this error meansh
	Message() string
}

type CommandInput map[string]any

type CommandError struct {
	Code         ErrorCode    `json:"code"`
	Message      string       `json:"message"`
	ProcessInput CommandInput `json:"processInput"`
	// TODO add native error
	// TODO helper message??
}

func (e *CommandError) Error() string {
	return e.Code.Message()
}

// NewExecutionError should this happend internally?
func NewExecutionError(code ErrorCode, input CommandInput) *CommandError {
	return &CommandError{
		Code:         code,
		ProcessInput: input,
		Message:      code.Message(),
	}
}

type QueryError string

func (q QueryError) Message() string {
	return "Error ExecutionQuery"
}

// NewExecutionError should this happend internally?
func NewQueryError(code ErrorCode, input CommandInput) *CommandError {
	return &CommandError{
		Code:         code,
		ProcessInput: input,
		Message:      code.Message(),
	}
}
