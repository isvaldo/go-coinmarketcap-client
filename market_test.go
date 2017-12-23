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

		convey.Convey("Response.name must be NXT", func() {
			response := TickerResponse{
				TickerList: []TickerItem{},
			}
			response.TickerList = append(response.TickerList, TickerItem{Name: "NXT"})
			response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin"})
			item, err := response.First()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "NXT")
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

		convey.Convey("Response.name must be Siacoin", func() {
			response := TickerResponse{
				TickerList: []TickerItem{},
			}
			response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin"})
			response.TickerList = append(response.TickerList, TickerItem{Name: "Litecoin"})
			response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin"})
			item, err := response.Last()
			convey.So(err, convey.ShouldBeNil)
			convey.So(item.Name, convey.ShouldEqual, "Siacoin")
		})
	})
}

func TestSortByPriceBtc(t *testing.T) {
	convey.Convey("SortByPriceBtc must be return list in order", t, func() {
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin",
			PriceBtc: "5"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin",
			PriceBtc: "10"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "NXT",
			PriceBtc: "2"})

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
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin",
			PercentChange1H: "5"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin",
			PercentChange1H: "10"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "NXT",
			PercentChange1H: "2"})

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
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin",
			PercentChange7D: "5"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin",
			PercentChange7D: "10"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "NXT",
			PercentChange7D: "2"})

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
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin",
			PercentChange24H: "5"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin",
			PercentChange24H: "10"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "NXT",
			PercentChange24H: "2"})

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
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin",
			AvailableSupply: "5"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin",
			AvailableSupply: "10"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "NXT",
			AvailableSupply: "2"})

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
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin",
			MarketCapUsd: "5"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin",
			MarketCapUsd: "10"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "NXT",
			MarketCapUsd: "2"})

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
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Name: "Siacoin"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "Bitcoin"})
		response.TickerList = append(response.TickerList, TickerItem{Name: "NXT"})

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
		response := TickerResponse{
			TickerList: []TickerItem{},
		}
		response.TickerList = append(response.TickerList, TickerItem{Symbol: "SC"})
		response.TickerList = append(response.TickerList, TickerItem{Symbol: "BTC"})
		response.TickerList = append(response.TickerList, TickerItem{Symbol: "NXT"})

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
