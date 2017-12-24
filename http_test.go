package coinmarket

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/smartystreets/goconvey/convey"
	"gopkg.in/jarcoal/httpmock.v1"
	"io/ioutil"
	"testing"
)

func getResponse(fileName string) (string, error) {
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("test/mocks/%s", fileName))
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("Error on open file %s", fileName))
	}
	return string(fileBytes), nil

}

func mustGetDummyTestServer(pathMock string, responseFile string) (Interface, error) {
	const url = "https://api.coinmarketcap.com"
	client := New(url)
	httpmock.Activate()
	mockResponse, err := getResponse(responseFile)
	clientPath := fmt.Sprintf("%s%s", url, pathMock)
	fmt.Println(clientPath)
	httpmock.RegisterResponder("GET", clientPath,
		httpmock.NewStringResponder(200, mockResponse))
	return client, err
}

func TestTickerClientMustBeWorker(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	convey.Convey("Given a dummy test server", t, func() {
		client, err := mustGetDummyTestServer("/v1/ticker/bitcoin/", "ticker.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
	})
}

func TestTickerMustReturnOneItem(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	convey.Convey("Given a dummy test server", t, func() {
		client, err := mustGetDummyTestServer("/v1/ticker/bitcoin/", "ticker.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)

		convey.Convey("Must return one item from ticker calls", func() {
			item, err := client.GetTicker("bitcoin")
			convey.So(err, convey.ShouldBeNil)
			convey.So(item, convey.ShouldNotBeNil)
		})
	})
}

func TestTickerShouldReturnBitcoinFields(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	convey.Convey("Given a dummy test server", t, func() {

		client, err := mustGetDummyTestServer("/v1/ticker/bitcoin/", "ticker.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
		convey.Convey("Must return mock values", func() {
			item, err := client.GetTicker("bitcoin")
			convey.So(err, convey.ShouldBeNil)

			convey.So(item.ID, convey.ShouldEqual, "bitcoin")
			convey.So(item.Name, convey.ShouldEqual, "Bitcoin")
			convey.So(item.Symbol, convey.ShouldEqual, "BTC")
			convey.So(item.Rank, convey.ShouldEqual, "1")
			convey.So(item.PriceUsd, convey.ShouldEqual, "19222.2")
			convey.So(item.PriceBtc, convey.ShouldEqual, "1.0")
			convey.So(item.Two4HVolumeUsd, convey.ShouldEqual, "13476500000.0")
			convey.So(item.MarketCapUsd, convey.ShouldEqual, "321917618660")
			convey.So(item.AvailableSupply, convey.ShouldEqual, "16748050.0")
			convey.So(item.MaxSupply, convey.ShouldEqual, "21000000.0")
			convey.So(item.PercentChange1H, convey.ShouldEqual, "-0.83")
			convey.So(item.PercentChange24H, convey.ShouldEqual, "-0.71")
			convey.So(item.PercentChange7D, convey.ShouldEqual, "22.8")
			convey.So(item.LastUpdated, convey.ShouldEqual, "1513545554")
		})
	})
}

func TestTickerLastMustBeReturnItemList(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	convey.Convey("Given a dummy test server", t, func() {

		client, err := mustGetDummyTestServer("/v1/ticker/", "tickerLast.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
		convey.Convey("When you call tickerLast your return must be list not empty", func() {

			response, err := client.GetTickerLast()
			convey.So(err, convey.ShouldBeNil)
			convey.So(response.TickerList, convey.ShouldNotBeEmpty)
		})
	})
}

func TestTickerLastMustBeReturnMockValues(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	convey.Convey("Given a dummy test server", t, func() {

		client, err := mustGetDummyTestServer("/v1/ticker/", "tickerLast.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
		convey.Convey("when you execute TestTickerLastMustBeReturnMockValues, the return must be mocked", func() {

			response, err := client.GetTickerLast()
			convey.So(err, convey.ShouldBeNil)
			coins := []string{"BTC", "ETH", "BCH", "XRP", "LTC", "ADA", "MIOTA",
				"DASH", "XEM", "XMR", "BTG", "XLM", "EOS", "NEO", "ETC",
			}
			for index, coin := range coins {
				convey.So(response.TickerList[index].Symbol, convey.ShouldEqual, coin)
			}
		})
	})
}

func TestTickerLimitMustBeLimitedResult(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	convey.Convey("Given a dummy test server", t, func() {
		client, err := mustGetDummyTestServer("/v1/ticker/?limit=3", "tickerWithLimit.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
		convey.Convey("when you apply limit for 3 itens in GetTickerWithLimits must be respected", func() {
			response, err := client.GetTickerWithLimits(3)
			convey.So(err, convey.ShouldBeNil)
			convey.So(response.Size(), convey.ShouldEqual, 3)
		})

		convey.Convey("when you apply limit for 3 itens in GetTickerWithLimits must return mock values", func() {
			response, err := client.GetTickerWithLimits(3)
			convey.So(err, convey.ShouldBeNil)
			convey.So(response.TickerList[0].Name, convey.ShouldEqual, "Bitcoin")
		})

	})
}

func TestTickerWithRangeMustBeReturnMockRange(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	convey.Convey("Given a dummy test server", t, func() {
		client, err := mustGetDummyTestServer("/v1/ticker/?start=0&limit=3", "tickerWithLimit.json")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
		convey.Convey("when you apply start 0 and limit 3 in GetTickerInRange must be respected", func() {
			response, err := client.GetTickerInRange(0, 3)
			convey.So(err, convey.ShouldBeNil)
			convey.So(response.Size(), convey.ShouldEqual, 3)
		})

		convey.Convey("when you apply start 0 and limit 3 in GetTickerInRange must return mock values", func() {
			response, err := client.GetTickerInRange(0, 3)
			convey.So(err, convey.ShouldBeNil)
			convey.So(response.TickerList[0].Name, convey.ShouldEqual, "Bitcoin")
		})

	})
}
