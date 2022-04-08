package handler

import (
	"email-template-generator/app/user"
	"email-template-generator/auth"
	"email-template-generator/entity"
	"email-template-generator/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{
		service:     service,
		authService: authService,
	}
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

	signedToken, err := h.authService.GenerateToken(response.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIFailedResponse(
			"Failed To Login",
			http.StatusBadRequest,
			gin.H{"errors": err.Error()},
		))
		return
	}

	response.Token = signedToken
	response, err = h.service.UpdateToken(response)

	c.JSON(http.StatusOK, helper.APIResponse(
		"Successfully Login",
		http.StatusOK,
		response,
	))
}

func (h *userHandler) ChangePassword(c *gin.Context) {
	var input user.ChangePasswordInput
	currentUser := c.MustGet("current_user").(entity.User)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := helper.ErrorMessageResponse(errors)

		c.JSON(http.StatusBadRequest, helper.APIFailedResponse(
			"Failed to Change Password",
			http.StatusBadRequest,
			errorMessage,
		))
		return
	}

	response, err := h.service.ChangePassword(currentUser, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIFailedResponse(
			"Failed to Change Password",
			http.StatusBadRequest,
			gin.H{"errors": err.Error()},
		))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponse(
		"Success Change Password",
		http.StatusOK,
		response,
	))
}
