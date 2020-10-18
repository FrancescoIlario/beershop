package rest

type ErrorCode int

const (
	ErrCodeValidationFailed = 1
	ErrCodeConflict         = 2
	ErrCodeInternal         = 3
)

type E struct {
	Code       ErrorCode         `json:"code"`
	Message    string            `json:"message"`
	Validation map[string]string `json:"validation"`
}
