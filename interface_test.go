package coinmarket

import (
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	convey.Convey("Must start client with 'New'", t, func() {
		client := New("https://api.coinmarketcap.com")

		convey.So(client, convey.ShouldNotBeNil)
	})
}

func TestNewCoinMarketCapWithCustomClient(t *testing.T) {
	convey.Convey("Must start client with 'NewWithClient' and custom client", t, func() {
		client := NewWithClient("https://api.coinmarketcap.com", http.DefaultClient)
		convey.So(client, convey.ShouldNotBeNil)
	})
}
