module github.com/go-echarts/examples

go 1.18

require github.com/go-echarts/go-echarts/v2 v2.4.0-rc2

require github.com/montanaflynn/stats v0.7.1

require (
	cloud.google.com/go v0.99.0 // indirect
	github.com/alpacahq/alpaca-trade-api-go v1.9.0 // indirect
	github.com/alpacahq/alpaca-trade-api-go/v2 v2.8.0 // indirect
	github.com/alpacahq/alpaca-trade-api-go/v3 v3.5.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
)

// dev mode
//replace github.com/go-echarts/go-echarts/v2 => ../../go-echarts
//
//replace github.com/go-echarts/snapshot-chromedp => ../snapshot-chromedp
