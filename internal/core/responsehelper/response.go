package responsehelper

type CommonResponse struct {
	Message string `json:"message"`
}

func NewCommonResponse(message string) *CommonResponse {
	return &CommonResponse{
		Message: message,
	}
}
