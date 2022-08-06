package v1

import "math"

type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
	PageSize   int64 `json:"size,omitempty"`
	Page       int64 `json:"page,omitempty"`
	Pages      int64 `json:"pages,omitempty"`
	FirstFlag  bool  `json:"firstFlag,omitempty"`
	LastFlag   bool  `json:"lastFlag,omitempty"`
}

type ListOptions struct {
	//TypeMeta      `json:",inline"`
	ListMeta       `json:"pageInfo,omitempty"`
	Unscoped       bool   `json:"unscoped"`
	LabelSelector  string `json:"labelSelector,omitempty" form:"labelSelector"`
	FieldSelector  string `json:"fieldSelector,omitempty" form:"fieldSelector"`
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
	Offset         *int64 `json:"offset,omitempty" form:"offset"`
	Limit          *int64 `json:"limit,omitempty" form:"limit"`
	PageFlag       bool   `json:"pageFlag,omitempty" form:"pageFlag"`
}

func (page *ListOptions) Init() *ListOptions {
	page.PageFlag = false
	page.FirstFlag = true
	if page.Page > 0 {
		// 当前
		page.Limit = &page.PageSize
		offset := (page.Page - 1) * page.PageSize
		page.Offset = &offset
		page.FirstFlag = false
		page.PageFlag = true
		if page.Page > 1 {
			page.FirstFlag = false
		}
	}
	return page
}

func (page *ListOptions) IsLast() *ListOptions {
	page.LastFlag = true
	if page.PageFlag {
		if page.Pages <= 0 {
			page.Pages = int64(math.Ceil(float64(page.TotalCount) / float64(page.PageSize)))
		}
		if page.Pages > page.Page {
			page.LastFlag = false
		}
	}
	return page
}
