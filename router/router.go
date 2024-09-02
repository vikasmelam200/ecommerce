package router

import (
	"onvif310824/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/cameras", controllers.AddCamera)
	r.GET("/cameras/:id/stream", controllers.StreamCamera)
	r.PUT("/cameras/:id/quality", controllers.SetImageQuality)
	r.PUT("/cameras/:id/bitrate", controllers.SetBitrate)
	r.PUT("/cameras/:id/timestamp", controllers.SetTimestamp)
	r.PUT("/cameras/:id/flip", controllers.SetFlip)
	r.PUT("/cameras/:id/mirror", controllers.SetMirror)
	r.PUT("/cameras/:id/overlay", controllers.SetTextOverlay)
	r.PUT("/cameras/:id/exposure", controllers.SetExposureControl)

	return r
}
