package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIError struct {
	StatusCode int
	Err        error
	Message    string
}

func (e *APIError) Error() string {
	return e.Err.Error()
}

func NewAPIError(statusCode int, err error, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Err:        err,
		Message:    message,
	}
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() //Only inside middleware. It executes the pending handlers in the chain inside the calling handler.

		// Check if there are any errors
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Check if it's an APIError
				var apiErr *APIError
				if errors.As(e.Err, &apiErr) {
					c.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
					return
				}
			}

			// If it's not an APIError, return a generic server error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
	}
}
