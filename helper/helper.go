package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, Code int, Status string, Data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    Code,
		Status:  Status,
	}

	response := Response{
		Meta: meta,
		Data: Data,
	}

	return response
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
