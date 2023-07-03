package helper

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func CustomError(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param())
				}

				break
			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
}

func FieldErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s tidak boleh kosong", e.StructField())
	case "email":
		return "Email tidak valid"
	}
	return ""
}

func ToResponse(c echo.Context, err error, data interface{}) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]string, len(ve))
		for i, fe := range ve {
			out[i] = FieldErrorToText(fe)
		}
		return c.JSON(http.StatusBadRequest, ResponseData{Status: http.StatusBadRequest, Error: true, Message: strings.Join(out, ", ")})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseData{Status: http.StatusBadRequest, Error: true, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseData{Status: http.StatusOK, Error: false, Message: "Success", Data: data})
}
