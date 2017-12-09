package coinmarket

type Interface interface {
	getTicker() error
	getTickerWithLimits(limit int)
	getTickerInRange(start int, end int)
}

