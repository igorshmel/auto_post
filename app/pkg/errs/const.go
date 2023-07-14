package errs

// Dictionary with errors
const (
	// Common errors
	Internal       = "ERR_INTERNAL"
	Syntax         = "ERR_SYNTAX"
	NotFound       = "ERR_NOT_FOUND"
	EmptyDbPointer = "ERR_EMPTY_DB_POINTER"
	EmptyInputData = "ERR_EMPTY_INPUT_DATA"

	// Доступ к ресурсам
	AccessDenied = "ERR_ACCESS_DENIED"

	// Validation
	ParseRequest      = "ERR_PARSE_REQUEST"
	ParametersInvalid = "ERR_PARAMETERS_INVALID"
)
