package router

import (
	"example.com/go-rest-api-gin/controllers/photoscontroller"
	"example.com/go-rest-api-gin/controllers/userscontroller"
	"example.com/go-rest-api-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.RouterGroup) {
	g.GET("/users", middlewares.Auth, userscontroller.GetUsers)
	g.GET("/users/:id", middlewares.Auth, userscontroller.GetUserById)
	g.POST("/users/register", userscontroller.AddUser)
	g.POST("/users/login", userscontroller.Login)
	g.PUT("/users/:id", middlewares.Auth, userscontroller.UpdateUser)
	g.DELETE("/users/:id", middlewares.Auth, userscontroller.DeleteUser)
}

func PhotoRoutes(g *gin.RouterGroup) {
	g.GET("/photos", middlewares.Auth, photoscontroller.GetPhotos)
	g.GET("/photos/:id", middlewares.Auth, photoscontroller.GetPhotoById)
	g.POST("/photos", middlewares.Auth, photoscontroller.AddPhoto)
	g.PUT("/photos/:id", middlewares.Auth, photoscontroller.UpdatePhoto)
	g.DELETE("/photos/:id", middlewares.Auth, photoscontroller.DeletePhoto)
}
