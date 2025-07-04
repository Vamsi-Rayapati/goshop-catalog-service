package apicommon

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
