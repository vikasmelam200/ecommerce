package controllers

import (
	"net/http"
	"onvif310824/models"
	"onvif310824/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddCamera(ctx *gin.Context) {
	var camera models.Camera
	err := ctx.ShouldBindJSON(&camera)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	services.AddCamera(&camera)
	ctx.JSON(200, camera)
}
func StreamCamera(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	streamURL, err := services.StreamCamera(cameraID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"stream url": streamURL})
}
func SetImageQuality(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	qualityStr := ctx.Query("quality")

	quality, err := strconv.Atoi(qualityStr)
	if err != nil || quality < 0 || quality > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image quality. Must be between 0 and 100."})
		return
	}

	err = services.SetImageQuality(cameraID, quality)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Image quality updated successfully"})
}
func SetBitrate(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	bitrateStr := ctx.Query("bitrate")

	bitrate, err := strconv.Atoi(bitrateStr)
	if err != nil || bitrate < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bitrate. Must be a positive integer."})
		return
	}

	err = services.SetBitrate(cameraID, bitrate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bitrate updated successfully"})
}
func SetTimestamp(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	enabledStr := ctx.Query("enabled")

	enabled, err := strconv.ParseBool(enabledStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for enabled. Must be true or false."})
		return
	}

	err = services.SetTimestamp(cameraID, enabled)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Timestamp setting updated successfully"})
}
func SetTextOverlay(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	overlayText := ctx.Query("text")

	if overlayText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Text overlay cannot be empty"})
		return
	}

	err := services.SetTextOverlay(cameraID, overlayText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Text overlay updated successfully"})
}
func SetFlip(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	flipStr := ctx.Query("flip")

	flip, err := strconv.ParseBool(flipStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for flip. Must be true or false."})
		return
	}

	err = services.SetFlip(cameraID, flip)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Flip setting updated successfully"})
}
func SetMirror(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	mirrorStr := ctx.Query("mirror")

	mirror, err := strconv.ParseBool(mirrorStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for mirror. Must be true or false."})
		return
	}

	err = services.SetMirror(cameraID, mirror)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mirror setting updated successfully"})
}
func SetExposureControl(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	exposureControlStr := ctx.Query("exposureControl")

	exposureControl, err := strconv.Atoi(exposureControlStr)
	if err != nil || exposureControl < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for exposure control. Must be a non-negative integer."})
		return
	}

	err = services.SetExposureControl(cameraID, exposureControl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Exposure control setting updated successfully"})
}
func SetGain(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	gainStr := ctx.Query("gain")

	gain, err := strconv.Atoi(gainStr)
	if err != nil || gain < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for gain. Must be a non-negative integer."})
		return
	}

	err = services.SetGain(cameraID, gain)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Gain setting updated successfully"})
}
func SetBacklightCompensation(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	backlightCompensationStr := ctx.Query("backlightCompensation")

	backlightCompensation, err := strconv.ParseBool(backlightCompensationStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for backlight compensation. Must be true or false."})
		return
	}

	err = services.SetBacklightCompensation(cameraID, backlightCompensation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Backlight compensation setting updated successfully"})
}
func SetPrivacyMasks(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	privacyMasks := ctx.Query("privacyMasks")

	if privacyMasks == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Privacy masks cannot be empty"})
		return
	}

	err := services.SetPrivacyMasks(cameraID, privacyMasks)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Privacy masks updated successfully"})
}
func SetScheduledProfileSettings(ctx *gin.Context) {
	cameraID := ctx.Param("id")
	scheduledProfileSettings := ctx.Query("scheduledProfileSettings")

	if scheduledProfileSettings == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Scheduled profile settings cannot be empty"})
		return
	}

	err := services.SetScheduledProfileSettings(cameraID, scheduledProfileSettings)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Scheduled profile settings updated successfully"})
}
