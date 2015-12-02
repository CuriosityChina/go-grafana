package grafana

import "strings"

type GetDashboardResponse struct {
	Dashboard Dashboard `json:"dashboard"`
	Meta      Meta      `json:"meta"`
}

type PostDashboardRequest struct {
	Dashboard Dashboard `json:"dashboard"`
	Overwrite bool      `json:"overwrite"`
}

type PostDashboardResponse struct {
	Slug    string `json:"slug,omitmepty"`
	Status  string `json:"status"`
	Version int64  `json:"version,omitempty"`
	message string `json:"message,omitempty"`
}

type Dashboard struct {
	Annotations struct {
		List []interface{} `json:"list"`
	} `json:"annotations"`
	Editable        bool        `json:"editable"`
	HideControls    bool        `json:"hideControls"`
	Id              int64       `json:"id"`
	Links           []Link      `json:"links"`
	OriginalTitle   string      `json:"originalTitle"`
	Refresh         interface{} `json:"refresh"`
	Rows            []Row       `json:"rows"`
	SchenmaVersion  int64       `json:"schemaVersion"`
	SharedCrosshair bool        `json:"shardCrosshair"`
	Style           string      `json:"style"`
	Tags            []string    `json:"tags"`
	Templating      struct {
		List []interface{} `jsno:"list"`
	} `json:"templating"`
	Time       Time       `json:"time"`
	Timepicker Timepicker `json:"timepicker"`
	Timezone   string     `json:"timezone"`
	Title      string     `json:"title"`
	Version    int64      `json:"version"`
}

func NewDefaultDashBoard() *Dashboard {
	d := new(Dashboard)
	d.Annotations.List = make([]interface{}, 0)
	d.Links = make([]Link, 0)
	d.Tags = make([]string, 0)
	d.Templating.List = make([]interface{}, 0)
	d.Style = "dark"
	d.Timezone = "browser"
	d.HideControls = false
	d.SharedCrosshair = false
	d.Editable = true
	d.Refresh = "30s"
	d.Time = NewTime("now-6h", "now")
	d.Timepicker = NewDefaultTimepicker()
	return d
}

func (d *Dashboard) SetTitle(title string) *Dashboard {
	d.Title = strings.TrimSpace(title)
	d.OriginalTitle = d.Title
	return d
}

func (d *Dashboard) SetStyle(style string) *Dashboard {
	d.Style = strings.TrimSpace(style)
	return d
}

func (d *Dashboard) SetTimezone(timezone string) *Dashboard {
	d.Style = strings.TrimSpace(timezone)
	return d
}

func (d *Dashboard) SetRefresh(refresh interface{}) *Dashboard {
	switch v := refresh.(type) {
	case bool:
		d.SetRefresh(v)
	case string:
		d.SetRefresh(v)
	default:
		d.SetRefresh(false)
	}
	return d
}

func (d *Dashboard) AddRow(row Row) *Dashboard {
	d.Rows = append(d.Rows, row)
	return d
}

func (d *Dashboard) AddTag(tag string) *Dashboard {
	d.Tags = append(d.Tags, strings.TrimSpace(tag))
	return d
}

func (d *Dashboard) SetTime(from, to string) *Dashboard {
	d.Time = NewTime(from, to)
	return d
}

//  func (d *Dashboard)

type Link struct {
	Icon string   `json:"icon"`
	Tags []string `json:"tags"`
	Type string   `json:"type"`
}

func NewLink(tag string) Link {
	var link Link
	link.Icon = "external link"
	link.Type = "dashboards"
	link.Tags = append(link.Tags, tag)
	return link
}

type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func NewTime(from, to string) Time {
	var t Time
	t.From = strings.TrimSpace(from)
	t.To = strings.TrimSpace(to)
	return t
}

type Timepicker struct {
	Now              bool     `json:"now"`
	RefreshIntervals []string `json:"refresh_intervals"`
	TimeOptions      []string `json:"time_options"`
}

func NewDefaultTimepicker() Timepicker {
	var tp Timepicker
	tp.Now = true
	tp.RefreshIntervals = []string{
		"5s",
		"10s",
		"30s",
		"1m",
		"5m",
		"15m",
		"30m",
		"1h",
		"2h",
		"1d",
	}
	tp.TimeOptions = []string{
		"5m",
		"15m",
		"1h",
		"6h",
		"12h",
		"24h",
		"2d",
		"7d",
		"30d",
	}
	return tp
}
