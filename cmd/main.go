package main

import (
	"fmt"
	"github.com/isvaldo/go-coinmarketcap-client"
	"strconv"
	"time"
)

var Client coinmarket.Interface

func init() {
	var err error
	Client, err = coinmarket.New("https://api.coinmarketcap.com")
	if err != nil {
		panic(err)
	}
}

type Signal struct {
	Ticker coinmarket.TickerItem
}

func getSignal(tickerSignal chan Signal, duration time.Duration) {

	for x := range time.Tick(duration) {
		ticker, err := Client.GetTickerLast()
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("--------%d------", x.Minute()))
		ticker.SortByPercentChange1H()
		for _, tickerItem := range ticker.TickerList {
			tickerSignal <- Signal{
				Ticker: tickerItem,
			}
		}
	}

}
func main() {
	signal := make(chan Signal)

	go getSignal(signal, time.Second*20)

	lasts := make(map[string][]string)

	for item := range signal {
		if len(lasts[item.Ticker.Name]) == 0 {
			fmt.Println(item.Ticker.Name, item.Ticker.PercentChange1H)
			lasts[item.Ticker.Name] = append(lasts[item.Ticker.Name], item.Ticker.PercentChange1H)

		} else {
			oldValue, _ := strconv.ParseFloat(lasts[item.Ticker.Name][len(lasts[item.Ticker.Name])-1], 64)
			newValue, _ := strconv.ParseFloat(item.Ticker.PercentChange1H, 64)

			if oldValue-newValue > 0 {

				fmt.Println(item.Ticker.Name, fmt.Sprintf("Change: %f %%", oldValue-newValue))
			}
		}

	}

}
