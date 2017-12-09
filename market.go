package coinmarket

import (
	"github.com/pkg/errors"
	"sort"
	"strconv"
	"strings"
)

//TickerItem Define a snapshot of the coin market value
type TickerItem struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

//TickerResponse response, store list of tickeritem
type TickerResponse struct {
	TickerList []TickerItem
}

//Size returns the total amount returned
func (c *TickerResponse) Size() int {
	return len(c.TickerList)
}

//First return first item of response list
func (c *TickerResponse) First() (*TickerItem, error) {
	if len(c.TickerList) == 0 {
		return nil, errors.New("response is empty")
	}
	return &c.TickerList[0], nil
}

//Last return last item from response list
func (c *TickerResponse) Last() (*TickerItem, error) {
	if len(c.TickerList) == 0 {
		return nil, errors.New("response is empty")
	}
	return &c.TickerList[len(c.TickerList)-1], nil
}

//SortByPriceBtc sort response by price of bicoin field
//ref field: PriceBtc
func (c *TickerResponse) SortByPriceBtc() {
	sort.Slice(c.TickerList, func(i, j int) bool {

		itemA, _ := strconv.ParseFloat(c.TickerList[i].PriceBtc, 64)
		itemB, _ := strconv.ParseFloat(c.TickerList[j].PriceBtc, 64)
		return itemA > itemB
	})
}

//SortByPercentChange1H sort response by percent of change in 1 hour
//ref field: PercentChange1H
func (c *TickerResponse) SortByPercentChange1H() {
	sort.Slice(c.TickerList, func(i, j int) bool {
		itemA, _ := strconv.ParseFloat(c.TickerList[i].PercentChange1H, 64)
		itemB, _ := strconv.ParseFloat(c.TickerList[j].PercentChange1H, 64)
		return itemA > itemB
	})
}

//PercentChange24H sort response by percent of change in 24 hour
//ref field: PercentChange24H
func (c *TickerResponse) SortByPercentChange24H() {
	sort.Slice(c.TickerList, func(i, j int) bool {
		itemA, _ := strconv.ParseFloat(c.TickerList[i].PercentChange24H, 64)
		itemB, _ := strconv.ParseFloat(c.TickerList[j].PercentChange24H, 64)
		return itemA > itemB
	})
}

//PercentChange7D sort response by percent of change in 7 days
//ref field: PercentChange7D
func (c *TickerResponse) SortByPercentChange7D() {
	sort.Slice(c.TickerList, func(i, j int) bool {
		itemA, _ := strconv.ParseFloat(c.TickerList[i].PercentChange7D, 64)
		itemB, _ := strconv.ParseFloat(c.TickerList[j].PercentChange7D, 64)
		return itemA > itemB
	})
}

//SortByAvailableSupply sort response by available supply value
//ref field: AvailableSupply
func (c *TickerResponse) SortByAvailableSupply() {
	sort.Slice(c.TickerList, func(i, j int) bool {
		itemA, _ := strconv.ParseFloat(c.TickerList[i].AvailableSupply, 64)
		itemB, _ := strconv.ParseFloat(c.TickerList[j].AvailableSupply, 64)
		return itemA > itemB
	})
}

//MarketCapUsd sort response by available supply value
//ref field: MarketCapUsd
func (c *TickerResponse) SortByMarketCapUsd() {
	sort.Slice(c.TickerList, func(i, j int) bool {
		itemA, _ := strconv.ParseFloat(c.TickerList[i].MarketCapUsd, 64)
		itemB, _ := strconv.ParseFloat(c.TickerList[j].MarketCapUsd, 64)
		return itemA > itemB
	})
}

//GetCoinNameStartsWith return list of tickerElements starting with prefix name
func (c *TickerResponse) GetCoinNameStartsWith(prefix string) []TickerItem {
	var response []TickerItem
	for index := range c.TickerList {
		if strings.HasPrefix(strings.ToLower(c.TickerList[index].Name), strings.ToLower(prefix)) {
			response = append(response, c.TickerList[index])
		}
	}
	return response
}

//GetCoinSymbolStartsWith return list of tickerElements starting with prefix symbol
func (c *TickerResponse) GetCoinSymbolStartsWith(prefix string) []TickerItem {
	var response []TickerItem
	for index := range c.TickerList {
		if strings.HasPrefix(strings.ToLower(c.TickerList[index].Symbol), strings.ToLower(prefix)) {
			response = append(response, c.TickerList[index])
		}
	}
	return response
}
