package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20002
	ErrInvalidToken     = 30001
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "uid is invalid",
	ErrInvalidToken:     "Invalid token",
}
