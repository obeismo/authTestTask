package handler

import (
	"medods/pkg/database"
	"medods/pkg/model"
	"medods/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Home page"})
}

func signUp(c *gin.Context) {
	var input model.User

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Errorf(err.Error())
		return
	}

	userRefreshToken := uuid.New()

	input.RefreshToken = service.GenerateHash(userRefreshToken.String())

	id := database.CreateUser(input)
	input.ID = id

	token, err := service.GenerateToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id":            id,
		"token":         token,
		"refresh_token": userRefreshToken,
	})
}

func signIn(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	refreshDB := database.CheckRefreshValid(input)
	if refreshDB == "" {
		logrus.Errorf("invalid refresh token")
		return
	}

	newToken, err := service.GenerateToken(input)
	if err != nil {
		logrus.Errorf("user not found: %s", err.Error())
		return
	}

	newRefreshToken := uuid.New()
	input.RefreshToken = service.GenerateHash(newRefreshToken.String())

	id, err := database.UpdateUser(input)
	if err != nil {
		logrus.Errorf("user not found: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":            id,
		"token":         newToken,
		"refresh_token": newRefreshToken,
	})
}
