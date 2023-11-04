package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type App struct {
	ID              int       `json:"id" 				gorm:"column:id;type:bigserial;primary_key;unique"`
	UID             uuid.UUID `json:"uid"   		    gorm:"column:uid;type:uuid;not null"`
	TeamID          int       `json:"teamID" 		    gorm:"column:team_id;type:bigserial"`
	Name            string    `json:"name" 				gorm:"column:name;type:varchar"`
	ReleaseVersion  int       `json:"releaseVersion" 	gorm:"column:release_version;type:bigserial"`
	MainlineVersion int       `json:"mainlineVersion" 	gorm:"column:mainline_version;type:bigserial"`
	Config          string    `json:"config" 	        gorm:"column:config;type:jsonb"`
	CreatedAt       time.Time `json:"createdAt" 		gorm:"column:created_at;type:timestamp"`
	CreatedBy       int       `json:"createdBy" 		gorm:"column:created_by;type:bigserial"`
	UpdatedAt       time.Time `json:"updatedAt" 		gorm:"column:updated_at;type:timestamp"`
	UpdatedBy       int       `json:"updatedBy" 		gorm:"column:updated_by;type:bigserial"`
	EditedBy        string    `json:"editedBy"          gorm:"column:edited_by;type:jsonb"`
}
