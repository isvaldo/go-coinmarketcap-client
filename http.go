package coinmarket

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTP struct {
	BaseURL    string
	HTTPClient *http.Client
}

const (
	contextPath         = "/v1/ticker"
	getTickerPath       = "%s/%s/"
	getTickerLast       = "%s/"
	getTickerWithLimits = "%s/?limit=%d"
	getTickerInRange    = "%s/?start=%d&limit=%d"
)

//getTicker resquest info of one coin (Ex bitcoin)
func (c *HTTP) getTicker(coinName string) (*TickerItem, error) {
	var responseRaw []json.RawMessage
	url := fmt.Sprintf("%s"+getTickerPath, c.BaseURL, contextPath, coinName)
	if err := c.doRequest("GET", url, nil, &responseRaw); err != nil {
		return nil, errors.Wrapf(err, "Error on perform request url:[%s]", url)
	}
	item := &TickerItem{}
	if err := json.Unmarshal(responseRaw[0], item); err != nil {
		return nil, errors.Wrapf(err, "Error on unmarshal data:[%s]", responseRaw[0])
	}
	return item, nil
}

func (c *HTTP) getTickerLast() (*TickerResponse, error) {
	var items []TickerItem
	url := fmt.Sprintf("%s"+getTickerLast, c.BaseURL, contextPath)
	if err := c.doRequest("GET", url, nil, &items); err != nil {
		return nil, errors.Wrapf(err, "Error on perform request url:[%s]", url)
	}
	return &TickerResponse{TickerList: items}, nil
}

func (c *HTTP) getTickerWithLimits(limit int) (*TickerResponse, error) {
	var items []TickerItem
	url := fmt.Sprintf("%s"+getTickerWithLimits, c.BaseURL, contextPath, limit)
	if err := c.doRequest("GET", url, nil, &items); err != nil {
		return nil, errors.Wrapf(err, "Error on perform request url:[%s]", url)
	}
	return &TickerResponse{TickerList: items}, nil
}

func (c *HTTP) getTickerInRange(start int, end int) (*TickerResponse, error) {
	var items []TickerItem
	url := fmt.Sprintf("%s"+getTickerInRange, c.BaseURL, contextPath, start, end)
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
