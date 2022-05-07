package main

import (
	"log"

	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/binance"
	"github.com/nntaoli-project/goex/builder"
)

// ref. https://dev.binance.vision/t/why-do-i-see-this-error-invalid-api-key-ip-or-permissions-for-action/93
const (
	SPOT_TESTNET_BASE_URL                = "https://testnet.binance.vision"
	SPOT_PRODUCTION_BASE_URL             = "https://api.binance.com"
	FUTURES_TESTNET_BASE_URL             = "https://testnet.binancefuture.com"
	FUTURES_PRODUCTION_BASE_URL          = "https://fapi.binance.com"
	DELIVERY_FUTURES_TESTNET_BASE_URL    = "https://testnet.binancefuture.com"
	DELIVERY_FUTURES_PRODUCTION_BASE_URL = "https://dapi.binance.com "
)

const (
	BINANCE_TESTNET_API_KEY        = "YOUR_KEY"
	BINANCE_TESTNET_API_KEY_SECRET = "YOUR_KEY_SECRET"
)

var binanceApi goex.FutureRestAPI

func initBuilder() {
	binanceApi = builder.DefaultAPIBuilder.APIKey(BINANCE_TESTNET_API_KEY).APISecretkey(BINANCE_TESTNET_API_KEY_SECRET).FuturesEndpoint(FUTURES_TESTNET_BASE_URL).BuildFuture(goex.BINANCE_SWAP)
}

func fetchFutureDepthAndIndex() {

	depth, err := binanceApi.GetFutureDepth(goex.BTC_USD, goex.SWAP_USDT_CONTRACT, 100)
	if err != nil {
		log.Fatalln(err.Error())
	}

	askTotalAmount, bidTotalAmount := 0.0, 0.0
	askTotalVol, bidTotalVol := 0.0, 0.0

	for _, v := range depth.AskList {
		askTotalAmount += v.Amount
		askTotalVol += v.Price * v.Amount
	}

	for _, v := range depth.BidList {
		bidTotalAmount += v.Amount
		bidTotalVol += v.Price * v.Amount
	}

	markPrice, err := binanceApi.GetFutureIndex(goex.BTC_USD)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Printf("CURRENT mark price: %f", markPrice)

	log.Printf("ContractType: %s ContractId: %s Pair: %s UTime: %s AmountTickSize: %d\n", depth.ContractType, depth.ContractId, depth.Pair, depth.UTime.String(), depth.Pair.AmountTickSize)
	log.Printf("askTotalAmount: %f, bidTotalAmount: %f, askTotalVol: %f, bidTotalVol: %f", askTotalAmount, bidTotalAmount, askTotalVol, bidTotalVol)
	log.Printf("ask price averge: %f, bid price averge: %f,", askTotalVol/askTotalAmount, bidTotalVol/bidTotalAmount)
	log.Printf("ask-bid spread: %f%%,", 100*(depth.AskList[0].Price-depth.BidList[0].Price)/markPrice)
}

func subscribeFutureMarketData() {
	binanceWs, err := builder.DefaultAPIBuilder.APIKey(BINANCE_TESTNET_API_KEY).APISecretkey(BINANCE_TESTNET_API_KEY_SECRET).Endpoint(binance.TESTNET_FUTURE_USD_WS_BASE_URL).BuildFuturesWs(goex.BINANCE_FUTURES)

	if err != nil {
		log.Fatalln(err.Error())
	}
	binanceWs.TickerCallback(func(ticker *goex.FutureTicker) {
		//log.Printf("%+v\n", *ticker.Ticker)
	})
	binanceWs.SubscribeTicker(goex.BTC_USD, goex.SWAP_USDT_CONTRACT)
	binanceWs.DepthCallback(func(depth *goex.Depth) {
		log.Printf("%+v\n", *depth)
	})
	binanceWs.SubscribeDepth(goex.BTC_USDT, goex.SWAP_USDT_CONTRACT)

	binanceWs.TradeCallback(func(trade *goex.Trade, contractType string) {
		log.Printf("%+v\n", *trade)
	})
	binanceWs.SubscribeTrade(goex.BTC_USDT, goex.SWAP_USDT_CONTRACT)

	select {}
}

func subscribeSpotMarketData() {

	binanceWs, err := builder.DefaultAPIBuilder.APIKey(BINANCE_TESTNET_API_KEY).APISecretkey(BINANCE_TESTNET_API_KEY_SECRET).Endpoint(binance.TESTNET_FUTURE_USD_BASE_URL).BuildSpotWs(goex.BINANCE)

	if err != nil {
		log.Fatalln(err.Error())
	}
	binanceWs.TickerCallback(func(ticker *goex.Ticker) {
		log.Printf("%+v\n", *ticker)
	})
	binanceWs.SubscribeTicker(goex.BTC_USDT)
	binanceWs.DepthCallback(func(depth *goex.Depth) {
		log.Printf("%+v\n", *depth)
	})
	binanceWs.SubscribeDepth(goex.BTC_USDT)

	binanceWs.TradeCallback(func(trade *goex.Trade) {
		log.Printf("%+v\n", *trade)
	})
	binanceWs.SubscribeTrade(goex.BTC_USDT)

	select {}

}

func main() {

	initBuilder()
	//subscribeFutureMarketData()
	//subscribeFutureMarketData()
	//fetchFutureDepthAndIndex()

	futuresOrder, err := binanceApi.MarketFuturesOrder(goex.BTC_USD, goex.SWAP_USDT_CONTRACT, "1.0", goex.OPEN_BUY, 10)

	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("%+v", futuresOrder)
}
