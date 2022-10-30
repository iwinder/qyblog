package v1

import (
	"encoding/json"
	"gorm.io/plugin/soft_delete"
	"time"
)

type ObjectMeta struct {
	ID           uint64                `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENt;column:id;comment:主键"`
	InstanceID   string                `json:"instanceID,omitempty" gorm:"unique;colum:instance_id;type:varchar(32);default:'def-0'';not null;comment:唯一资源标识"`
	Sort         int                   `json:"Sort" gorm:"colum:sort;default:0"`
	StatusFlag   int                   `json:"status,omitempty" gorm:"column:status_flag;default:1;comment:启用标识，1 启用，2 禁用" validate:"omitempty"`
	Extend       Extend                `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ExtendShadow string                `json:"-" gorm:"column:extend_shadow" validate:"omitempty"`
	createdBy    uint64                `json:"-" gorm:"column:created_by;default:1;comment:创建者" `
	CreatedAt    *time.Time            `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt    *time.Time            `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DeletedAt    *time.Time            `json:"-" gorm:"index;comment:删除时间"`
	DeletedFlag  soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0;"`
}

type ObjectNoInstMeta struct {
	ID           uint64     `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENt;column:id;comment:主键"`
	InstanceID   string     `json:"instanceID,omitempty" gorm:"colum:instance_id;type:varchar(32);default:'def-0'';comment:唯一资源标识"`
	Sort         int        `json:"Sort" gorm:"colum:sort;default:0"`
	StatusFlag   int        `json:"status,omitempty" gorm:"column:status_flag;default:0;comment:启用标识，0 启用，1 禁用,2 删除" validate:"omitempty"`
	Extend       Extend     `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ExtendShadow string     `json:"-" gorm:"column:extend_shadow" validate:"omitempty"`
	CreatedAt    *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
	DisabledAt   *time.Time `json:"-,omitempty" gorm:"column:disabled_at"`
	DeletedAt    *time.Time `json:"-" gorm:"index;comment:删除时间"`
}

func (ext Extend) String() string {
	data, _ := json.Marshal(ext)
	return string(data)
}

func (ext Extend) Merge(extendShadow string) Extend {
	var extend Extend
	_ = json.Unmarshal([]byte(extendShadow), &extend)
	for k, v := range extend {
		if _, ok := ext[k]; !ok {
			ext[k] = v
		}
	}
	return ext
}
