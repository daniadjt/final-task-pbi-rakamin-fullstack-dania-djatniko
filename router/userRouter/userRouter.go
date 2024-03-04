package userRouter

import (
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/models"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/controllers/usercontroller"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/helpers"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnvVariables()

	models.ConnectDatabase()
}

func UserRouter(r *gin.Engine) {

	r.GET("/users", usercontroller.Index)
	r.GET("/user/:id", usercontroller.Show)
	r.POST("/user/register", usercontroller.Register)
	r.POST("/user/login", usercontroller.Login)
	r.PUT("/user/:id", usercontroller.Update)
	r.DELETE("/user/:id", usercontroller.Delete)
}