package grafana

type Config struct {
	URL   string
	Token string
}

func NewConfig(url, token string) Config {
	return Config{
		URL:   url,
		Token: token,
	}
}
