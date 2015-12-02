package grafana

import "strings"

type Row struct {
	Collapse  bool    `json:"collapse"`
	Editable  bool    `json:"editable"`
	Height    string  `json:"height"`
	Panels    []Panel `json:"panels"`
	ShowTitle bool    `json:"showTitle"`
	Title     string  `json:"title"`
}

func NewDefaultRow(title string) Row {
	var row Row
	row.Collapse = false
	row.Editable = true
	row.Height = "200px"
	row.Title = strings.TrimSpace(title)
	row.ShowTitle = true
	return row
}

func (r *Row) SetTitle(title string) *Row {
	r.Title = strings.TrimSpace(title)
	return r
}

func (r *Row) SetHeight(h string) *Row {
	r.Height = strings.TrimSpace(h)
	return r
}

func (r *Row) AddPanel(p Panel) *Row {
	r.Panels = append(r.Panels, p)
	return r
}
