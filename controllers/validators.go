package controllers

import (
	"fmt"
	"go-identity-service/models"
	"go-identity-service/services"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
)

func ValidateUser(user *models.User) (string, int) {
	v := validator.New()
	var userService = services.UserService{}
	_ = v.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		_, err := userService.GetOne(strings.ToLower(fl.FieldName()), fl.Field().String())
		return err != nil
	})

	validationErrors := v.Struct(user)
	if validationErrors != nil {
		for _, e := range validationErrors.(validator.ValidationErrors) {
			fmt.Println(e)
			return fmt.Sprintf("Error in %v field", e.Field()), http.StatusBadRequest
		}
	}
	return "", 200
}
