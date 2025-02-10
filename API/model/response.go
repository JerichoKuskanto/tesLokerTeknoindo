package model

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Generic Response
type SuccessResponseModelFactory interface {
	CreateSuccessResponse(message string, data interface{}) *Response
}

type ConcreteSuccessResponseModelFactory struct{}

// Membuat response untuk request yang berhasil (status 200)
func (factory *ConcreteSuccessResponseModelFactory) CreateSuccessResponse(message string, data interface{}) *Response {
	return &Response{
		Status:  200,
		Message: message,
		Data:    data,
	}
}

// Membuat response model factory baru
func NewSuccessResponseModelFactory() SuccessResponseModelFactory {
	return &ConcreteSuccessResponseModelFactory{}
}

// Error Response
type ErrorResponseModelFactory interface {
	CreateErrorResponse(message string) *Response
}

type ConcreteErrorResponseModelFactory struct{}

// Membuat response untuk request yang gagal (status 400)
func (factory *ConcreteErrorResponseModelFactory) CreateErrorResponse(message string) *Response {
	return &Response{
		Status:  400,
		Message: message,
		Data:    nil,
	}
}

func NewErrorResponseModelFactory() ErrorResponseModelFactory {
	return &ConcreteErrorResponseModelFactory{}
}
