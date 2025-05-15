package wrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

type CustomError struct {
	Message        string      `json:"message"`
	StatusCode     int         `json:"statusCode"`
	HttpStatusCode int         `json:"httpStatusCode"`
	Data           interface{} `json:"data"`
}

func (ce CustomError) Error() string {
	return ce.Message
}

func SendSuccessResponse(ctx *gin.Context, message string, data interface{}, code int) {
	response := CustomResponse{
		Success: true,
		Message: message,
		Data:    data,
		Code:    code,
	}

	ctx.JSON(http.StatusOK, response)
}

func SendErrorResponse(ctx *gin.Context, err error, data interface{}, code int) {
	statusCode := code
	httpCode := code
	responseData := data

	if customErr, ok := err.(CustomError); ok {
		statusCode = customErr.StatusCode
		httpCode = customErr.HttpStatusCode
		if customErr.Data != nil {
			responseData = customErr.Data
		}
	}

	response := CustomResponse{
		Success: false,
		Message: err.Error(),
		Data:    responseData,
		Code:    statusCode,
	}

	ctx.JSON(httpCode, response)
}