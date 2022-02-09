package v1

import (
	"encoding/json"
	"time"
)

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}
type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

type ListOptions struct {
	TypeMeta `json:",inline"`

	LabelSelector string `json:"labelSelector,omitempty" form:"labelSelector"`

	FieldSelector string `json:"fieldSelector,omitempty" form:"fieldSelector"`

	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`

	Offset *int64 `json:"offset,omitempty" form:"offset"`

	Limit *int64 `json:"limit,omitempty" form:"limit"`
}
type ObjectMeta struct {
	ID         uint64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENt;column:id"`
	InstanceID string `json:"instanceID,omitempty" gorm:"unique;colum:instance_id;type:varchar(32);not null"`
	//Name       string    `json:"name,omitempty" gorm:"colum:name;type:varchar(255);not null"`
	Extend       Extend    `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ExtendShadow string    `json:"-" gorm:"column:extend_shadow" validate:"omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}
type CreateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}

type GetOptions struct {
	TypeMeta `json:",inlime"`
}
type Extend map[string]interface{}

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
