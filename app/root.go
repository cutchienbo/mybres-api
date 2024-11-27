package app

import (
	"fmt"
	"os"
	"sample/app/controllers"
	"sample/app/helper"
	"sample/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitRoot() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	helper.GormDB = gORMConnection()

	r.GET("/user/jwt", controllers.TestJWT)       

	r.GET("/get-access-token", controllers.GetAccessToken)

	r.POST("/user/signin", controllers.UserSignInController)     

	r.GET("/user/list", middleware.AuthGuard, controllers.UserGetController)       
	r.POST("/user/entry", controllers.UserEntryController)   
	r.POST("/user/delete", controllers.UserDeleteController) 

	return r
}

func gORMConnection() *gorm.DB {
	var db *gorm.DB

	var (
		devHostName = os.Getenv("MYSQL_HOST")
		devDbName   = os.Getenv("MYSQL_DB_NAME")
		devUser     = os.Getenv("MYSQL_USER")
		devPassword = os.Getenv("MYSQL_PASSWORD")
		devPort     = os.Getenv("MYSQL_PORT")
	)

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", devUser, devPassword, devHostName, devPort, devDbName)+"?parseTime=true&charset=utf8mb4&loc=Local"),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)

	if err != nil {
		panic(err)
	}
	
	return db
}
