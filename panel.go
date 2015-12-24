package grafana

type GraphPanel struct {
	AliasColors     interface{}   `json:"aliasColors,omitempty"`
	Bars            bool          `json:"bars"`
	Datasource      string        `json:"datasource,omitempty"`
	Editable        bool          `json:"editable"`
	Error           bool          `json:"error"`
	Fill            int64         `json:"fill"`
	Grid            interface{}   `json:"grid,omitempty"`
	Id              int64         `json:"id"`
	LeftYAxisLabel  string        `json:"leftYAxisLabel"`
	Legend          Legend        `json:"legend,omitempty"`
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

func (l *Legend) EnableAvg() *Legend {
	l.Show = true
	l.Values = true
	l.Avg = true
	return l
}

func (l *Legend) DisableAvg() *Legend {
	l.Avg = false
	return l
}

func (l *Legend) EnableCurrent() *Legend {
	l.Show = true
	l.Values = true
	l.Current = true
	return l
}

func (l *Legend) DisableCurrent() *Legend {
	l.Avg = false
	return l
}

func (l *Legend) EnableMax() *Legend {
	l.Show = true
	l.Values = true
	l.Max = true
	return l
}

func (l *Legend) DisableMax() *Legend {
	l.Max = false
	return l
}

func (l *Legend) EnableMin() *Legend {
	l.Show = true
	l.Values = true
	l.Min = true
	return l
}

func (l *Legend) DisableMin() *Legend {
	l.Min = false
	return l

}

func (l *Legend) EnableTotal() *Legend {
	l.Show = true
	l.Values = true
	l.Total = true
	return l

}

func (l *Legend) DisableTotal() *Legend {
	l.Total = false
	return l
}

func (l *Legend) EnableLegend() *Legend {
	l.Show = true
	l.Values = true
	return l
}

func (l *Legend) DisableLegend() *Legend {
	l.Show = false
	l.Values = true
	return l
}

func NewGraphPanel(title string) GraphPanel {
	var p GraphPanel
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
	p.Legend.EnableLegend()
	return p
}

func (p *GraphPanel) SetDatasource(ds string) *GraphPanel {
	p.Datasource = ds
	return p
}

func (p *GraphPanel) SetStack(check bool) *GraphPanel {
	p.Stack = check
	return p
}

func (p *GraphPanel) AddTarget(target Target) *GraphPanel {
	p.Targets = append(p.Targets, target)
	return p
}

func (p *GraphPanel) SetSpanWidth(width int) *GraphPanel {
	if width > 12 || width < 0 {
		width = 4
	}
	p.Span = width
	return p
}
