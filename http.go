package coinmarket

import (
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
)

type HTTP struct {
	BaseURL string
	HTTPClient *http.Client
}

func (c *HTTP) getTicker(coin string) (*TickerResponse, error) {

	return nil,nil
}

func (c *HTTP)  getTickerLast() (*TickerResponse, error){
	return nil,nil
}

func (c *HTTP) getTickerWithLimits(limit int) (*TickerResponse, error) {
	return nil,nil
}

func (c *HTTP) getTickerInRange(start int, end int) (*TickerResponse, error){
	return nil,nil
}

//doRequest is help for execute request and parse the response
func (c *HTTP) doRequest(method, url string, body io.Reader, to interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	//start := time.Now()
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if to == nil {
		io.Copy(ioutil.Discard, res.Body)
		return nil
	}
	//elapsed := time.Since(start)
	return json.NewDecoder(res.Body).Decode(to)
}
