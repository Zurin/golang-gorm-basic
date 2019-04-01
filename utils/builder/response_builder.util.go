package builder

//response message
const (
	MessageFetchTrxFailed  = "Failed to get transactions data"
	MessageFetchTrxSuccess = "Transactions data retrieved successfully"
	MessageAuthFailed      = "Authentication failed"
)

type ErrResponse struct {
	Error string
}

type MsgResponse struct {
	Message string
}

type ReadyStatement struct {
	Dbconn int
	Query  string
	Params []interface{}
}

func BaseResponse(success bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": success,
		"message": message,
		"data":    data}
}

func WebsocketResponse(success bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": success,
		"message": message,
		"data":    data}
}
