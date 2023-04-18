package errs

// Error constants
const (
	MsgEmptyDbPointer = "database pointer is empty"
	MsgEmptyInputData = "input data is empty"
	MsgNotFound       = "not found"

	UnknownError = "ERR_UNKNOWN_ERROR"
)

// CommonErrorResponse is response
type CommonErrorResponse struct {
	Domain  string `json:"domain"`  // Наименование сервиса (обязательное поле)
	Code    int    `json:"code"`    // Код ошибки в рамках сервиса (обязательное поле)
	Reason  string `json:"reason"`  // Описание кода ошибки
	Context string `json:"context"` // Дополнительная информация, позволяющая определить проблему
	KbLink  string `json:"kb_link"` // Ссылка на ресурс, описывающий решение проблемы
}

// NewCommonErrorResponse is constructor
func NewCommonErrorResponse(code int, reason string, context string, kbLink string) *CommonErrorResponse {
	return &CommonErrorResponse{
		Domain:  "dfa_auto_post/",
		Code:    code,
		Reason:  reason,
		Context: context,
		KbLink:  kbLink,
	}
}
