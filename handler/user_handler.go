package handler

import (
	"golang-CRUD/helper"
	"golang-CRUD/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) FindAll(c *gin.Context) {
	findUser, err := h.service.FindAll()
	if err != nil {
		response := helper.APIResponse("User is not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User found", http.StatusOK, "success", findUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) FindByID(c *gin.Context) {
	userID := c.Param("ID")
	id, _ := strconv.Atoi(userID)

	findUser, err := h.service.FindByID(id)
	if err != nil {
		response := helper.APIResponse("User is not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User found", http.StatusOK, "success", findUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Create(c *gin.Context) {
	var userInput user.UserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create user failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.service.Create(userInput)

	if err != nil {
		response := helper.APIResponse("Create user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Create user success", http.StatusOK, "success", newUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Update(c *gin.Context) {
	var userInput user.UserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create user failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := c.Param("ID")
	id, _ := strconv.Atoi(userID)

	updateUser, err := h.service.Update(id, userInput)
	if err != nil {
		response := helper.APIResponse("Update user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update user succeeded", http.StatusOK, "success", updateUser)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Delete(c *gin.Context) {
	userID := c.Param("ID")
	id, _ := strconv.Atoi(userID)

	deleteUser, err := h.service.Delete(id)
	if err != nil {
		response := helper.APIResponse("Delete user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete user succeeded", http.StatusOK, "success", deleteUser)
	c.JSON(http.StatusOK, response)
}
