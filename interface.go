package coinmarket

import "net/http"

type Interface interface {
	GetTicker(coinName string) (*TickerItem, error)
	GetTickerLast() (*TickerResponse, error)
	GetTickerWithLimits(limit int) (*TickerResponse, error)
	GetTickerInRange(start int, end int) (*TickerResponse, error)
}

//NewWithClient return a instance from MarketTicker, with optional client configuration
func NewWithClient(baseURL string, client *http.Client) Interface {
	return NewMarketTicker(baseURL, client)
}

//NewWithClient Normal call
func New(baseURL string) Interface {
	return NewMarketTicker(baseURL, http.DefaultClient)
}

//NewMedia return a instance from MarketTicker, with optional client configuration
func NewMarketTicker(baseURL string, client *http.Client) Interface {
	return &HTTP{
		BaseURL:    baseURL,
		HTTPClient: client,
	}
}
