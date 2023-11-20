package models

import "github.com/m3rashid/awesome/db"

const DRIVE_FILE_MODEL_NAME = "drivefiles"

type DriveFile struct {
	db.BaseModel
	Name        string `json:"name" gorm:"column:name;not null" validate:"required"`
	Type        string `json:"type" gorm:"column:type;not null" validate:"required"`
	Parent      uint   `json:"parent" gorm:"column:parent;not null" validate:"required"`
	ResourceUrl string `json:"resourceUrl" gorm:"column:resourceUrl;not null" validate:"required"`
	IsFolder    bool   `json:"isFolder" gorm:"column:isFolder;not null" validate:"required"`
}

func (DriveFile) TableName() string {
	return DRIVE_FILE_MODEL_NAME
}