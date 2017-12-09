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
	response, err := client.getTickerInRange(0, 100)

	if err != nil {
		panic(err)
	}
	response.SortByAvailableSupply()
	for index := range response.TickerList {
		fmt.Println(response.TickerList[index].Name, response.TickerList[index].AvailableSupply)
	}
}
