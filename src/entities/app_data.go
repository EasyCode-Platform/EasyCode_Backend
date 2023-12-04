package entities

import (
	"github.com/google/uuid"
	"time"
)

type App struct {
	ID          int       `json:"id" 				gorm:"column:id;type:bigserial;primary_key;unique"`
	UID         uuid.UUID `json:"uid"   		    gorm:"column:uid;type:uuid;not null"`
	TeamID      int       `json:"teamID" 		    gorm:"column:team_id"`
	Name        string    `json:"name" 				gorm:"column:name;type:text"`
	ComponentId string    `json:"name" 				gorm:"column:coponent_id;type:text;notnull"`
	Config      string    `json:"config" 	        gorm:"column:config;type:jsonb"`
	CreatedAt   time.Time `json:"createdAt" 		gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `json:"updatedAt" 		gorm:"column:updated_at;type:timestamp"`
}

// App 表示应用的数据结构
type AppData struct {
	Aid    uuid.UUID `json:"aid"`
	Name   string    `json:"name"`
	Tables []Table   `json:"tables"`
}

type Table struct {
	Tid     uuid.UUID `json:"tid"`
	Name    string    `json:"name"`
	AppAid  uuid.UUID `json:"-"`       // 序列化时忽略此字段
	Fields  []Field   `json:"fields"`  // 新增字段，用于存储字段信息
	Records []Record  `json:"records"` // 新增字段，用于存储记录信息
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Record map[string]interface{}
