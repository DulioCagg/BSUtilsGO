package resterr

import "net/http"

// RestErr is used to stablish a common error response throughout the application
type RestErr struct {
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Type    string   `json:"type"`
	Causes  []string `json:"causes"`
}

// BadRequest creates a RestErr with the given message
func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Type:    "bad_request",
	}
}

// NotFound creates a RestErr with the given message
func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Type:    "not_found",
	}
}

// InternalServerError creates a RestErr with the given message
func InternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Type:    "internal_server_error",
	}

	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}
