package photoRouter

import (
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/models"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/controllers/photocontroller"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/helpers"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnvVariables()

	models.ConnectDatabase()

}
func PhotoRouter(r *gin.Engine) {

	r.GET("/photos", helpers.Auth, photocontroller.GetAllPhoto)
	r.GET("/photo/:id", helpers.Auth, photocontroller.GetPhotoById)
	r.POST("/photo", helpers.Auth, photocontroller.Create)
	r.PUT("/photo/:id", helpers.Auth, photocontroller.Update)
	r.DELETE("/photo/:id", helpers.Auth, photocontroller.Delete)

}