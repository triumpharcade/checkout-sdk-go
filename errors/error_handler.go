package errors

import (
	"encoding/json"
	"net/http"
)

func HandleError(statusCode int, status string, requestId string, body []byte) CheckoutAPIError {
	if statusCode >= http.StatusInternalServerError {
		return CheckoutAPIError{
			StatusCode: statusCode,
			Status:     string(body),
		}
	}

	var details ErrorDetails
	if len(body) != 0 {
		if err := json.Unmarshal(body, &details); err != nil {
			details.RequestID = requestId

			return CheckoutAPIError{
				StatusCode: http.StatusBadRequest,
				Status:     string(body),
				Data:       &details,
			}
		}
	}

	details.RequestID = requestId

	return CheckoutAPIError{
		StatusCode: statusCode,
		Status:     status,
		Data:       &details,
	}
}
