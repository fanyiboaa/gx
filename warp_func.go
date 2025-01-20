package gx

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ErrorHandler   func(*gin.Context, error)
	SuccessHandler func(*gin.Context, any)
	BindFunc       func(*gin.Context, any) error
)

var (
	errorHandler   ErrorHandler   = defaultErrorHandler
	successHandler SuccessHandler = defaultSuceessHandler
	bindFunc       BindFunc       = defaultBindFunc
)

func SetErrorHandler(h ErrorHandler) {
	errorHandler = h
}

func SetSuccessHandler(h SuccessHandler) {
	successHandler = h
}

func SetBindFunc(f BindFunc) {
	bindFunc = f
}

func defaultErrorHandler(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
}

func defaultSuceessHandler(ctx *gin.Context, result any) {
	ctx.JSON(http.StatusOK, result)
}

func defaultBindFunc(ctx *gin.Context, obj any) error {
	if err := ctx.ShouldBindJSON(obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error binding JSON: %v", err)})
		return err
	}
	return nil
}

func W[T any](f func(*T) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req T
		if err := bindFunc(ctx, &req); err != nil {
			return
		}

		result, err := f(&req)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		successHandler(ctx, result)
	}
}

func WE(f func() (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := f()
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		successHandler(ctx, result)
	}
}

func WC[T any](f func(*gin.Context, *T) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req T
		if err := bindFunc(ctx, &req); err != nil {
			return
		}

		result, err := f(ctx, &req)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		successHandler(ctx, result)
	}
}

func WCE(f func(*gin.Context) (any, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := f(ctx)
		if err != nil {
			errorHandler(ctx, err)
			return
		}

		successHandler(ctx, result)
	}
}
