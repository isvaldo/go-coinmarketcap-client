package coinmarket

import (
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

func TestMustClient(t *testing.T) {
	convey.Convey("Must start client with 'must' method no fail allowed", t, func() {
		client := Must("https://api.coinmarketcap.com")
		convey.So(client, convey.ShouldNotBeNil)
	})
}

func TestNewClient(t *testing.T) {
	convey.Convey("Must start client with 'New'", t, func() {
		client, err := New("https://api.coinmarketcap.com")
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
	})
}

func TestNewCoinMarketCapWithCustomClient(t *testing.T) {
	convey.Convey("Must start client with 'NewWithClient' and custom client", t, func() {
		client, err := NewWithClient("https://api.coinmarketcap.com", http.DefaultClient)
		convey.So(err, convey.ShouldBeNil)
		convey.So(client, convey.ShouldNotBeNil)
	})
}
