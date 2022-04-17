package v1

type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

type ListOptions struct {
	TypeMeta      `json:",inline"`
	Unscoped      bool   `json:"unscoped"`
	LabelSelector string `json:"labelSelector,omitempty" form:"labelSelector"`

	FieldSelector string `json:"fieldSelector,omitempty" form:"fieldSelector"`

	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`

	Offset *int64 `json:"offset,omitempty" form:"offset"`

	Limit *int64 `json:"limit,omitempty" form:"limit"`
}
