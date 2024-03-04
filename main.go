package main

import (
	// "github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/seed"

	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/router/photoRouter"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/router/userRouter"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/helpers"
	"github.com/daniadjt/task-5-pbi-btpns-dania-djatniko/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	models.ConnectDatabase()
	userRouter.UserRouter(r)
	photoRouter.PhotoRouter(r)

	// seed.Load(models.DB)

	r.Run()
}
