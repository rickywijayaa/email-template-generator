package handler

import (
	"email-template-generator/app/user"
	"email-template-generator/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var user user.LoginInput
	err := c.ShouldBindJSON(&user)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := helper.ErrorMessageResponse(errors)

		c.JSON(http.StatusUnprocessableEntity, helper.APIFailedResponse(
			"Failed To Login",
			http.StatusUnprocessableEntity,
			errorMessage,
		))
		return
	}

	response, err := h.service.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIFailedResponse(
			"Failed To Login",
			http.StatusBadRequest,
			gin.H{"errors": err.Error()},
		))
		return
	}

	//JWT
	// token, err := h.AuthService.GenerateToken(loggedInUser.ID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, helper.APIFailedResponse(
	// 		"Failed To Login",
	// 		http.StatusBadRequest,
	// 		gin.H{"errors": err.Error()},
	// 	))
	// 	return
	// }

	c.JSON(http.StatusOK, helper.APIResponse(
		"Successfully Login",
		http.StatusOK,
		response,
	))
}
