package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"salarykart/ctypes"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "Should have length >= " + fe.Param()
	case "max":
		return "Should have length <= " + fe.Param()
	}
	if s, ok := ctypes.GetErrorMessage(fe); ok {
		return s
	}
	return "Unknown error"
}

func setError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		//TODO: make it set the content type as application/json
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}

	// handle JSON unmarshall errors
	var je *json.UnmarshalTypeError
	if errors.As(err, &je) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": []ErrorMsg{{
			Field:   je.Field,
			Message: je.Value,
		}}})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}
