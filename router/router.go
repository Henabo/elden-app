package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hiro942/elden-app/controller"
	"github.com/hiro942/elden-app/middleware"
)

// Routers initialize the main router
func Routers() *gin.Engine {
	Router := gin.Default()

	// config the access control
	Router.Use(middleware.Cors())

	NodeGroup := Router.Group("node")
	{
		NodeGroup.POST("/initLedger", controller.InitLedger)
		NodeGroup.POST("/satellite/register", controller.SatelliteRegister)
		NodeGroup.POST("/user/register", controller.UserRegister)
		NodeGroup.POST("/user/accessRecord", controller.CreateAccessRecord)
		NodeGroup.POST("/user/changeAuthStatus", controller.ChangeAuthStatus)
		NodeGroup.POST("delete", controller.DeleteNodeById)
		NodeGroup.GET("/satellite/publicKey", controller.GetSatellitePublicKey)
		NodeGroup.GET("/user/publicKey", controller.GetUserPublicKey)
		NodeGroup.GET("/search", controller.GetNodeById)
		NodeGroup.GET("/search/all", controller.GetAllNodes)
	}

	return Router
}
