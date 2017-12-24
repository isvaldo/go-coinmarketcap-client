package coinmarket

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCheckFirstMethod(t *testing.T) {
	convey.Convey("When first method of tickerResponse called empty", t, func() {

		convey.Convey("Response must be error", func() {
			response := TickerResponse{}
			item, err := response.First()
			convey.So(item, convey.ShouldBeNil)
			convey.So(err, convey.ShouldBeError)
		})

		convey.Convey("Response.name must be Siacoin", func() {
			response := getFakeResponse()
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Siacoin")
		})
	})
}

func TestCheckLastMethod(t *testing.T) {
	convey.Convey("When first method of tickerResponse called empty, the response must be error", t, func() {
		convey.Convey("Response must be error", func() {
			response := TickerResponse{}
			item, err := response.Last()
			convey.So(item, convey.ShouldBeNil)
			convey.So(err, convey.ShouldBeError)
		})

		convey.Convey("Response.name must be NXT", func() {
			response := getFakeResponse()
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
		})
	})
}

func TestSortByPriceBtc(t *testing.T) {
	convey.Convey("SortByPriceBtc must be return list in order", t, func() {
		response := getFakeResponse()
		convey.Convey("First item must be bitcoin", func() {
			response.SortByPriceBtc()
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Bitcoin")
		})

		convey.Convey("Last item must be NXT", func() {
			response.SortByPriceBtc()
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
		})

	})

}

func TestSortByPercentChange1H(t *testing.T) {
	convey.Convey("PercentChange1H must be return list in order", t, func() {
		response := getFakeResponse()
		convey.Convey("First item must be bitcoin", func() {
			response.SortByPercentChange1H()
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Bitcoin")
		})

		convey.Convey("Last item must be NXT", func() {
			response.SortByPercentChange1H()
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
		})

	})

}

func TestSortByPercentChange7D(t *testing.T) {
	convey.Convey("SortByPercentChange7D must be return list in order", t, func() {
		response := getFakeResponse()
		convey.Convey("First item must be bitcoin", func() {
			response.SortByPercentChange7D()
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Bitcoin")
		})

		convey.Convey("Last item must be NXT", func() {
			response.SortByPercentChange7D()
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
		})
	})
}
func TestSortByPercentChange24H(t *testing.T) {
	convey.Convey("SortByPercentChange24H must be return list in order", t, func() {
		response := getFakeResponse()
		convey.Convey("First item must be bitcoin", func() {
			response.SortByPercentChange24H()
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Bitcoin")
		})

		convey.Convey("Last item must be NXT", func() {
			response.SortByPercentChange24H()
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
		})
	})
}

func TestSortByAvailableSupply(t *testing.T) {
	convey.Convey("SortByAvailableSupply must be return list in order", t, func() {
		response := getFakeResponse()
		convey.Convey("First item must be bitcoin", func() {
			response.SortByAvailableSupply()
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Bitcoin")
		})

		convey.Convey("Last item must be NXT", func() {
			response.SortByAvailableSupply()
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
		})
	})
}

func TestSortByMarketCapUsd(t *testing.T) {
	convey.Convey("SortByMarketCapUsd must be return list in order", t, func() {
		response := getFakeResponse()
		convey.Convey("First item must be bitcoin", func() {
			response.SortByMarketCapUsd()
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Bitcoin")
		})

		convey.Convey("Last item must be NXT", func() {
			response.SortByMarketCapUsd()
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
		})
	})
}

func TestGetCoinNameStartWith(t *testing.T) {
	convey.Convey("Search by name prefix", t, func() {
		response := getFakeResponse()
		convey.Convey("Must be return Bitcoin", func() {
			Itens := response.GetCoinNameStartsWith("bit")
			convey.So(Itens[0].Name, convey.ShouldEqual, "Bitcoin")
		})

		convey.Convey("Must be return one item", func() {
			Itens := response.GetCoinNameStartsWith("bit")
			convey.So(len(Itens), convey.ShouldEqual, 1)
		})

	})
}

func TestGetCoinSymbolStartsWith(t *testing.T) {
	convey.Convey("Search by Symbol prefix", t, func() {
		response := getFakeResponse()
		convey.Convey("Must be return Bitcoin", func() {
			Itens := response.GetCoinSymbolStartsWith("btc")
			convey.So(Itens[0].Symbol, convey.ShouldEqual, "BTC")
		})

		convey.Convey("Must be return one item", func() {
			Itens := response.GetCoinSymbolStartsWith("btc")
			convey.So(len(Itens), convey.ShouldEqual, 1)
		})
	})
}

func getFakeResponse() *TickerResponse {
	response := TickerResponse{
		TickerList: []TickerItem{},
	}
	response.TickerList = append(response.TickerList, TickerItem{
		Symbol: "SC", Name: "Siacoin", MarketCapUsd: "5",
		PriceBtc: "5", AvailableSupply: "5", PercentChange24H: "5",
		PercentChange7D: "5", PercentChange1H: "5",
	})
	response.TickerList = append(response.TickerList, TickerItem{
		Symbol: "BTC", Name: "Bitcoin", MarketCapUsd: "10",
		PriceBtc: "10", AvailableSupply: "10", PercentChange24H: "10",
		PercentChange7D: "10", PercentChange1H: "10",
	})
	response.TickerList = append(response.TickerList, TickerItem{
		Symbol: "NXT", Name: "NXT", MarketCapUsd: "2",
		PriceBtc: "2", AvailableSupply: "2", PercentChange24H: "2",
		PercentChange7D: "2", PercentChange1H: "2",
	})
	return &response
}
