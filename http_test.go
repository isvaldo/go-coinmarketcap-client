package coinmarket

import (
	"fmt"
	"testing"
)

func TestTickerMustBeReturnOneItem(t *testing.T) {
	client, err := New("https://api.coinmarketcap.com")
	if err != nil {
		panic(err)
	}
	response, err := client.getTickerLast()

	if err != nil {
		panic(err)
	}

	for index := range response.TickerList {
		fmt.Println(response.TickerList[index].Name, response.TickerList[index].PercentChange1H)
	}
}
