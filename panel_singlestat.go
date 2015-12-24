package grafana

type SingleStatPanel struct {
	CacheTimeout    interface{}   `json:"cacheTimeout,omitempty"`
	ColorBackground bool          `json:"colorBackground,omitempty"`
	ColorValue      bool          `json:"colorValue,omitempty"`
	Colors          []string      `json:"colors,omitempty"`
	Datasource      interface{}   `json:"datasource"`
	Editable        bool          `json:"editable"`
	Error           bool          `json:"error"`
	Format          string        `json:"format"`
	ID              int           `json:"id"`
	Interval        interface{}   `json:"interval,omitempty"`
	IsNew           bool          `json:"isNew,omitempty"`
	Links           []interface{} `json:"links,omitempty"`
	MaxDataPoints   int           `json:"maxDataPoints"`
	NullPointMode   string        `json:"nullPointMode"`
	NullText        interface{}   `json:"nullText"`
	Postfix         string        `json:"postfix"`
	PostfixFontSize string        `json:"postfixFontSize"`
	Prefix          string        `json:"prefix"`
	PrefixFontSize  string        `json:"prefixFontSize"`
	Span            int           `json:"span"`
	Sparkline       struct {
		FillColor string `json:"fillColor"`
		Full      bool   `json:"full"`
		LineColor string `json:"lineColor"`
		Show      bool   `json:"show"`
	} `json:"sparkline"`
	Targets       []Target   `json:"targets"`
	Thresholds    string     `json:"thresholds"`
	Title         string     `json:"title"`
	Type          string     `json:"type"`
	ValueFontSize string     `json:"valueFontSize"`
	ValueMaps     []ValueMap `json:"valueMaps"`
	ValueName     string     `json:"valueName"`
}

type ValueMap struct {
	Op    string `json:"op"`
	Text  string `json:"text"`
	Value string `json:"value"`
}

func NewDefaultValueMap() ValueMap {
	var vm ValueMap
	vm.Op = "="
	vm.Text = "N/A"
	vm.Value = "null"
	return vm
}

func NewSingleStatPanel(title string) SingleStatPanel {
	var p SingleStatPanel
	p.Type = "signlestat"
	vm := NewDefaultValueMap()
	p.ValueMaps = append(p.ValueMaps, vm)
	p.Title = title
	return p
}

func (p *SingleStatPanel) SetDataSource(ds string) *SingleStatPanel {
	p.Datasource = ds
	return p
}

func (p *SingleStatPanel) AddTarget(target Target) *SingleStatPanel {
	p.Targets = append(p.Targets, target)
	return p
}

func (p *SingleStatPanel) SetSpanWidth(width int) *SingleStatPanel {
	if width > 12 || width < 0 {
		width = 4
	}
	p.Span = width
	return p
}
