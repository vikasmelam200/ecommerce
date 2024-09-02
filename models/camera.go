package models

import "time"

type Camera struct {
	ID                       uint   `gorm:"primaryKey"`
	IPAddress                string `gorm:"not null"`
	Username                 string
	Password                 string
	ImageWidth               int
	ImageHeight              int
	ImageQuality             int
	Bitrate                  int
	Timestamp                bool
	TextOverlay              string
	Flip                     bool
	Mirror                   bool
	Brightness               int
	Contrast                 int
	Saturation               int
	Sharpness                int
	WhiteBalance             int
	ExposureControl          int
	Gain                     int
	BacklightCompensation    bool
	PrivacyMasks             string
	ScheduledProfileSettings string
	CreatedAt                time.Time
}

type StreamLog struct {
	ID         uint   `gorm:"primaryKey"`
	CameraID   uint   `gorm:"not null"`
	StreamURL  string `gorm:"not null"`
	Resolution string
	Quality    int
	Bitrate    int
	Zoom       int
	StartedAt  time.Time
	EndedAt    *time.Time
}
