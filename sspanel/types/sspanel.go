package types

const (
	StatusFailure int = 0
	StatusOk          = 1
)

type ApiResponse struct {
	StatusCode int    `json:"ret"`
	Msg        string `json:"msg"`
}

type ReadNodeInfoResponse struct {
	ApiResponse
	Node NodeModal `json:"node"`
}
