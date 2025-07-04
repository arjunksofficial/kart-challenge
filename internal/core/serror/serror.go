package serror

type ServiceError struct {
	Code  int   `json:"code"`
	Error error `json:"error"`
}
