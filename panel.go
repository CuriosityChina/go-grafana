package grafana

type Panel struct {
	AliasColors     interface{}   `json:"aliasColors,omitempty"`
	Bars            bool          `json:"bars"`
	Datasource      string        `json:"datasource,omitempty"`
	Editable        bool          `json:"editable"`
	Error           bool          `json:"error"`
	Fill            int64         `json:"fill"`
	Grid            interface{}   `json:"grid,omitempty"`
	Id              int64         `json:"id"`
	LeftYAxisLabel  string        `json:"leftYAxisLabel"`
	Legend          interface{}   `json:"legend,omitempty"`
	Lines           bool          `json:"lines"`
	Linewidth       int64         `json:"linewidth"`
	Links           []interface{} `json:"links,omitempty"`
	MinSpan         interface{}   `json:"minSpan,omitempty"`
	NullPointMode   string        `json:"nullPointMode"`
	Percentage      bool          `json:"percentage"`
	Pointradius     int64         `json:"pointradius"`
	Points          bool          `json:"points"`
	Renderer        string        `json:"renderer"`
	RightYAxisLabel string        `json:"rightYAxisLabel"`
	SeriesOverrides []interface{} `json:"seriesOverrides,omitempty"`
	Span            interface{}   `json:"span,omitempty"`
	Stack           bool          `json:"stack"`
	SteppedLine     bool          `json:"steppedLine"`
	Targets         []Target      `json:"targets"`
	TimeFrom        interface{}   `json:"timeFrom,omitempty"`
	TimeShift       interface{}   `json:"timeShift,omitempty"`
	Title           string        `json:"title"`
	Tooltip         struct {
		Shared    bool   `json:"shared"`
		ValueType string `json:"value_type"`
	} `json:"tooltip"`
	Transparent bool     `json:"transparent"`
	Type        string   `json:"type"`
	XAxis       bool     `json:"x-axis"`
	YAxis       bool     `json:"y-axis"`
	YFormats    []string `json:"y_formats"`
}

type Legend struct {
	Avg     bool `json:"avg"`
	Current bool `json:"current"`
	Max     bool `json:"max"`
	Min     bool `json:"min"`
	Show    bool `json:"show"`
	Total   bool `json:"total"`
	Values  bool `json:"values"`
}

func NewGraphPanel(title string) Panel {
	var p Panel
	p.Type = "graph"
	p.XAxis = true
	p.YAxis = true
	p.Lines = true
	p.Fill = 1
	p.Linewidth = 2
	p.Renderer = "flot"
	p.NullPointMode = "connected"
	p.YFormats = []string{
		"short",
		"short",
	}
	p.Title = title
	return p
}

func (p *Panel) SetDatasource(ds string) *Panel {
	p.Datasource = ds
	return p
}

func (p *Panel) SetStack(check bool) *Panel {
	p.Stack = check
	return p
}

func (p *Panel) AddTarget(target Target) *Panel {
	p.Targets = append(p.Targets, target)
	return p
}

func (p *Panel) SetSpanWidth(width int) *Panel {
	p.Span = width
	return p
}
