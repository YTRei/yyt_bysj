package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}


var (
	ErrorRequestBodyParseFailed = Err{Error: "Request body in not correct", ErrorCode: "001"}
	ErrorNotAuthUser = Err{Error: "User authentication failed.", ErrorCode: "002"}
	ErrorDBError = Err{Error: "DB ops failed", ErrorCode: "003"}
	ErrorInternalFaults = Err{Error: "Internal service error", ErrorCode: "004"}
)