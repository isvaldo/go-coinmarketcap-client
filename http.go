package coinmarket

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

// concrete Implementation of coinmarket.Interface
type HTTP struct {
	BaseURL    string
	HTTPClient *http.Client
}

//resources path's
const (
	ContextPath         = "/v1/ticker"
	GetTickerPath       = "%s/%s/"
	GetTickerLast       = "%s/"
	GetTickerWithLimits = "%s/?limit=%d"
	GetTickerInRange    = "%s/?start=%d&limit=%d"
)

//getTicker request the snapshot from finance market for a coin
func (c *HTTP) GetTicker(coinName string) (*TickerItem, error) {
	var responseRaw []json.RawMessage
	url := fmt.Sprintf("%s"+GetTickerPath, c.BaseURL, ContextPath, coinName)
	if err := c.doRequest("GET", url, nil, &responseRaw); err != nil {
		return nil, errors.Wrapf(err, "Error on perform request url:[%s]", url)
	}
	item := &TickerItem{}
	if err := json.Unmarshal(responseRaw[0], item); err != nil {
		return nil, errors.Wrapf(err, "Error on unmarshal data:[%s]", responseRaw[0])
	}
	return item, nil
}

//getTickerLast request last snapshot's from finance market of the last 100 ranked coins
func (c *HTTP) GetTickerLast() (*TickerResponse, error) {
	var items []TickerItem
	url := fmt.Sprintf("%s"+GetTickerLast, c.BaseURL, ContextPath)
	if err := c.doRequest("GET", url, nil, &items); err != nil {
		return nil, errors.Wrapf(err, "Error on perform request url:[%s]", url)
	}
	return &TickerResponse{TickerList: items}, nil
}

//getTickerLast request last snapshot's from finance market of the last ranked coins(with limit)
func (c *HTTP) GetTickerWithLimits(limit int) (*TickerResponse, error) {
	var items []TickerItem
	url := fmt.Sprintf("%s"+GetTickerWithLimits, c.BaseURL, ContextPath, limit)
	if err := c.doRequest("GET", url, nil, &items); err != nil {
		return nil, errors.Wrapf(err, "Error on perform request url:[%s]", url)
	}
	return &TickerResponse{TickerList: items}, nil
}

//getTickerInRange request a range of finance market snapshot
func (c *HTTP) GetTickerInRange(start int, end int) (*TickerResponse, error) {
	var items []TickerItem
	url := fmt.Sprintf("%s"+GetTickerInRange, c.BaseURL, ContextPath, start, end)
	if err := c.doRequest("GET", url, nil, &items); err != nil {
		return nil, errors.Wrapf(err, "Error on perform request url:[%s]", url)
	}
	return &TickerResponse{TickerList: items}, nil
}

//doRequest is help for execute request and parse the response
func (c *HTTP) doRequest(method, url string, body io.Reader, to interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if to == nil {
		io.Copy(ioutil.Discard, res.Body)
		return nil
	}
	return json.NewDecoder(res.Body).Decode(to)
}
