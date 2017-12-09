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

	startsWithFilter := response.GetCoinSymbolStartsWith("bt")
	for index := range startsWithFilter {
		fmt.Println(startsWithFilter[index].Symbol)
	}
}
