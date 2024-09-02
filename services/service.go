package services

import (
	"errors"
	"fmt"
	"onvif310824/config"
	"onvif310824/models"
	"time"

	//"github.com/deepch/onvif"
	"github.com/videonext/onvif"
	//"github.com/videonext/onvif/xsd/onvif"
)

func AddCamera(camera *models.Camera) error {
	db := config.GetDB()

	if camera.IPAddress == "" || camera.Username == "" || camera.Password == "" {
		return errors.New("camera IP address, username, and password must be provided")
	}

	var existingCamera models.Camera
	if err := db.Where("id = ?", camera.ID).First(&existingCamera).Error; err == nil {
		return errors.New("camera with this ID already exists")
	}

	if err := db.Create(camera).Error; err != nil {
		return err
	}

	return nil
}

func StreamCamera(cameraID string) (string, error) {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return "", err
	}

	client := onvif.NewClient(camera.IPAddress, camera.Username, camera.Password)

	profiles, err := client.GetMediaProfiles()
	if err != nil {
		return "", err
	}

	profile := profiles[0]
	encoderConfig, err := client.GetVideoEncoderConfiguration(profile.Token)
	if err != nil {
		return "", err
	}

	encoderConfig.RateControl.FrameRateLimit = 30

	if camera.ImageWidth > 0 && camera.ImageHeight > 0 {
		encoderConfig.Resolution.Width = camera.ImageWidth
		encoderConfig.Resolution.Height = camera.ImageHeight
	}
	if camera.ImageQuality > 0 && camera.ImageQuality <= 100 {
		encoderConfig.Quality = float64(camera.ImageQuality) / 100.0
	}
	if camera.Bitrate > 0 {
		encoderConfig.RateControl.BitrateLimit = camera.Bitrate
	}

	err = client.SetVideoEncoderConfiguration(encoderConfig)
	if err != nil {
		return "", err
	}

	streamURI, err := client.GetStreamUri(profile.Token)
	if err != nil {
		return "", err
	}

	log := models.StreamLog{
		CameraID:   camera.ID,
		StreamURL:  streamURI.Uri,
		Resolution: fmt.Sprintf("%dx%d", encoderConfig.Resolution.Width, encoderConfig.Resolution.Height),
		Quality:    camera.ImageQuality,
		Bitrate:    camera.Bitrate,
		Zoom:       1,
		StartedAt:  time.Now(),
	}
	if err := db.Create(&log).Error; err != nil {
		return "", err
	}

	return streamURI.Uri, nil
}
func SetImageQuality(cameraID string, quality int) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	if quality < 0 || quality > 100 {
		return errors.New("image quality must be between 0 and 100")
	}

	camera.ImageQuality = quality
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}
func SetBitrate(cameraID string, bitrate int) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	if bitrate < 0 {
		return errors.New("bitrate must be a positive integer")
	}

	camera.Bitrate = bitrate
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}

// SetTimestamp enables or disables timestamp for a camera.
func SetTimestamp(cameraID string, enabled bool) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	camera.Timestamp = enabled
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}
func SetTextOverlay(cameraID string, overlayText string) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	camera.TextOverlay = overlayText
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}

func SetFlip(cameraID string, flip bool) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	camera.Flip = flip
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}
func SetMirror(cameraID string, mirror bool) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	camera.Mirror = mirror
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}
func SetExposureControl(cameraID string, exposureControl int) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	if exposureControl < 0 {
		return errors.New("exposure control must be a non-negative integer")
	}

	camera.ExposureControl = exposureControl
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}
func SetGain(cameraID string, gain int) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	// Validate gain setting
	if gain < 0 {
		return errors.New("gain must be a non-negative integer")
	}

	camera.Gain = gain
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}
func SetBacklightCompensation(cameraID string, backlightCompensation bool) error {
	db := config.GetDB()

	var camera models.Camera
	if err := db.First(&camera, "id = ?", cameraID).Error; err != nil {
		return err
	}

	camera.BacklightCompensation = backlightCompensation
	if err := db.Save(&camera).Error; err != nil {
		return err
	}

	return nil
}
