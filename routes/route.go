package routes

import (
	"cart/app/controllers/UserController"
	"github.com/gin-gonic/gin"
)
func UserRoutes(router *gin.Engine){
	user:=router.Group("/users")
	user.GET("/",UserController.Index)
}

