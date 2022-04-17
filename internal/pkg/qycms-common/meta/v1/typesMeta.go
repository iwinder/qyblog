package v1

type Extend map[string]interface{}

type TypeMeta struct {
	Kind       string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
}

type CreateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}
type UpdateOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
}
type GetOptions struct {
	TypeMeta `json:",inline"`
	DryRun   []string `json:"dryRun,omitempty"`
	Unscoped bool     `json:"unscoped"`
}

type DeleteOptions struct {
	TypeMeta `json:",inline"`
	Unscoped bool `json:"unscoped"`
}
