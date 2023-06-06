package errs

// Error constants
const (
	MsgEmptyDbPointer = "ERR_DATABASE_POINTER_IS_EMPTY"
	MsgEmptyInputData = "ERR_INPUT_DATA_IS_EMPTY"
	MsgNotFound       = "ERR_NOT_FOUND"

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
		Domain:  "lic_auto_post/",
		Code:    code,
		Reason:  reason,
		Context: context,
		KbLink:  kbLink,
	}
}
