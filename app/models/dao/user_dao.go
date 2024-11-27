package dao

import (
	// "fmt"
	"sample/app/helper"
	"sample/app/models/db"
	"sample/app/models/request"
	"sample/app/models/response"
)

func UserCheckEmailExist(email *string) int {
	var user db.User

	if result := helper.GormDB.Debug().Where("email = ?", email).First(&user); result.RowsAffected == 0 {
		return 0
	}

	return 1
}

func UserCheckExist(req *request.UserSignInRequest) *db.User {
	var user db.User

	if result := helper.GormDB.Debug().Where("email = ? AND password = ?", req.Email, req.Password).First(&user); result.RowsAffected == 0 {
		return nil
	}

	return &user
}

func UserExec() response.UserDetailsResponse {
	var users []db.User

	helper.GormDB.Debug().Find(&users)

	return response.UserDetailsResponse{
		Users: users,
		// Des:      users.Password,
	}
}

func UserEntryExec(req request.UserEntryRequest) error {
	err := helper.GormDB.Debug().Create(&db.User{
		Id:       req.Id,
		Name:     req.UserName,
		Password: req.Des,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func UserDeleteExec(req request.UserDeleteRequest) error {

	err := helper.GormDB.Debug().Table("users").Where("id = ?", req.Id).Delete(&db.User{}).Error
	if err != nil {
		return err
	}
	return nil

}
