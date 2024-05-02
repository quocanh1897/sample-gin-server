package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
	apperror "github.com/quocanh1897/sample-gin-server/internal/error"
	dto "github.com/quocanh1897/sample-gin-server/internal/model/dto/response"
)

func ErrorHandler() gin.HandlerFunc {
	return ErrorHandlerMiddleware{}.HandlerFunc
}

type ErrorHandlerMiddleware struct{}

func (m ErrorHandlerMiddleware) HandlerFunc(ctx *gin.Context) {
	ctx.Next()

	err := ctx.Errors.Last()
	if err == nil {
		return
	}
	var finalErr apperror.Error

	switch apiError := err.Err.(type) {
	case *http.MaxBytesError:
		finalErr = apperror.NewRequestTooLargeError(apiError.Error())
	case apperror.Error:
		finalErr = apiError
	case validator.ValidationErrors:
		finalErr = apperror.NewInvalidFieldError("Invalid input")
	default:
		finalErr = apperror.NewServerError("Internal Server Error")
	}

	requestId := ctx.GetString(constant.RequestIdAttributeKey)
	finalErr.RequestId = requestId

	ctx.JSON(finalErr.StatusCode, dto.FailedResponse(finalErr))
}
