package grafana

type Meta struct {
	Type    string `json:"type"`
	CanSave bool   `json:"canSave"`
	CanEdit bool   `json:"canEdit"`
	CanStar bool   `json:"canStar"`
	Slug    string `json:"slug"`
	Expires string `json:"expires"`
	Created string `json:"Created"`
}
