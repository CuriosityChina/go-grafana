package grafana

type Target struct {
	Expr           string `json:"expr"`
	Interval       int    `json:"interval"`
	IntervalFactor int32  `json:"intervalFactor"`
	LegendFormat   string `json:"legendFormat"`
	Metric         string `json:"metric"`
	RefId          string `json:"refId"`
}

func NewTarget(expr, legendFormat, metric, refID string) Target {
	var t Target
	t.IntervalFactor = 1
	t.Expr = expr
	t.LegendFormat = legendFormat
	t.Metric = metric
	t.RefId = refID
	return t
}
