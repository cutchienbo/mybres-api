package controllers

import (
	"net/http"
	"sample/app/helper"
	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	var refreshToken string = c.GetHeader("RefreshToken")

	if err := helper.CheckJWT(refreshToken, "refresh"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": err.Error(),
				"status": http.StatusBadRequest,
			},
		})

		return
	}	

	var refreshTokenSub = helper.GetRefreshTokenSub(refreshToken)

	var accessToken string = helper.GenerateAccessToken(refreshTokenSub)

	c.JSON(
		http.StatusOK, gin.H{
			"accessToken": accessToken,
			"message": "Get accessToken success",
			"status": http.StatusOK,
		},
	)

	return
}