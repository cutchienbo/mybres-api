package controllers

import (
	"net/http"
	"sample/app/helper"
	"sample/app/models/dao"
	"sample/app/models/db"
	"sample/app/models/request"

	"github.com/gin-gonic/gin"
)

func TestJWT(c *gin.Context){
	var jwt string = c.DefaultQuery("jwt", "")

	c.JSON(http.StatusOK, helper.CheckJWT(jwt, "access"))

	return
}

func UserSignInController(c *gin.Context) {
	var req request.UserSignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": err.Error(),
					"code":    http.StatusBadRequest,
				},
			},
		)

		return
	}

	if emailExisted := dao.UserCheckEmailExist(&req.Email); emailExisted == 0 {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": "Email not found",
					"code":    http.StatusBadRequest,
				},
			},
		)

		return
	}

	var userExisted *db.User

	if userExisted = dao.UserCheckExist(&req); userExisted == nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": "Wrong password",
					"code":    http.StatusBadRequest,
				},
			},
		)

		return
	}

	var user helper.UserJWTSubject = helper.UserJWTSubject{
		Id: 	userExisted.Id,
		Name: 	userExisted.Name,
	}

	c.JSON(
		http.StatusOK, gin.H{
			"accessToken": helper.GenerateAccessToken(user),
			"refreshToken": helper.GenerateRefreshToken(user),
			"message": "Signin success",
			"status": http.StatusOK,
		},
	)

	return
}

func UserGetController(c *gin.Context) {
	//var req request.ChampionListRequest
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.JSON(
	//		http.StatusBadRequest, gin.H{
	//			"errors": gin.H{
	//				"message": err.Error(),
	//				"code":    http.StatusBadRequest,
	//			},
	//		},
	//	)
	//	return
	//}
	//
	//helper.ZapLog.Sugar().Debugf("Body request is: %v", req)
	data := dao.UserExec()

	c.JSON(http.StatusOK, data)

	return
}

func UserEntryController(c *gin.Context) {
	var req request.UserEntryRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": err.Error(),
					"code":    http.StatusBadRequest,
				},
			},
		)

		return
	}

	//helper.ZapLog.Sugar().Debugf("Body request is: %v", req)
	/// 400
	err := dao.UserEntryExec(req)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": err.Error(),
					"code":    http.StatusBadRequest,
				},
			},
		)

		return
	}

	// 500 internal
	// 400
	// 200

	c.JSON(http.StatusOK, gin.H{
		"message": "OKMEN",
	})

	return
}

func UserDeleteController(c *gin.Context) {
	var req request.UserDeleteRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": err.Error(),
					"code":    http.StatusBadRequest,
				},
			},
		)

		return
	}

	//helper.ZapLog.Sugar().Debugf("Body request is: %v", req)
	/// 400
	err := dao.UserDeleteExec(req)

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": err.Error(),
					"code":    http.StatusBadRequest,
				},
			},
		)

		return
	}

	// 500 internal
	// 400
	// 200

	c.JSON(http.StatusOK, gin.H{
		"message": "OKMEN",
	})
	
	return
}
