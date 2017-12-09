package coinmarket

import "net/http"

type Interface interface {
	getTicker(coinName string) (*TickerItem, error)
	getTickerLast() (*TickerResponse, error)
	getTickerWithLimits(limit int) (*TickerResponse, error)
	getTickerInRange(start int, end int) (*TickerResponse, error)
}

//NewWithClient return a instance from MarketTicker, with optional client configuration
func NewWithClient(baseURL string, client *http.Client) (Interface, error) {
	return NewMarketTicker(baseURL, client)
}

func Must(baseURL string) Interface {
	intf, err := New(baseURL)
	if err != nil {
		panic(err)
	}
	return intf
}

func New(baseURL string) (Interface, error) {
	return NewMarketTicker(baseURL, http.DefaultClient)
}

//NewMedia return a instance from MarketTicker, with optional client configuration
func NewMarketTicker(baseURL string, client *http.Client) (Interface, error) {
	return &HTTP{
		BaseURL:    baseURL,
		HTTPClient: client,
	}, nil
}
