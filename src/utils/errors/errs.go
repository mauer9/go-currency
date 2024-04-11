package errors

import "fmt"

// ErrSystem
const ErrSystem = "TestTask"

// Error structure
type ArgError struct {
	System           string `json:"system"`
	Status           int    `json:"status"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

// Error to formatted string
func (e *ArgError) Error() string {
	return fmt.Sprintf("%d %s", e.Status, e.DeveloperMessage)
}

// Set developer message and return
func (e *ArgError) SetDevMessage(developMessage string) *ArgError {
	e.DeveloperMessage = developMessage
	return e
}

// 5xx errors variables
var (
	InternalServerError = &ArgError{ErrSystem, 503, "internal server error", "internal server error"}
	DBConnectError      = &ArgError{ErrSystem, 503, "db connect error", "db connect error"}
	DBReadError         = &ArgError{ErrSystem, 503, "db read error", "db read error"}
	DBWriteError        = &ArgError{ErrSystem, 503, "db write error", "db write error"}
	ENVReadError        = &ArgError{ErrSystem, 503, "env variables reading error", "env variables reading error"}
	FilesystemReadError = &ArgError{ErrSystem, 503, "filesystem reading error", "filesystem reading error"}
	RPCError            = &ArgError{ErrSystem, 503, "RPC error", "RPC error"}
	SerializeError      = &ArgError{ErrSystem, 503, "serialization error", "serialization error"}
)

// 4xx errors variables
var (
	AccessDenied     = &ArgError{System: ErrSystem, Status: 403, Message: "access denied", DeveloperMessage: "access denied"}
	InvalidCharacter = &ArgError{ErrSystem, 400, "incorrect input", "incorrect input"}
	IncorrectRequest = &ArgError{ErrSystem, 400, "incorrect request", "incorrect request"}
	NotFound         = &ArgError{ErrSystem, 404, "not found", "not found"}
)
