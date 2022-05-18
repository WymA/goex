package bybit

import (
	"net/http"
	"testing"

	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/internal/logger"
)

var baDapi = NewByBitFutures(&goex.APIConfig{
	HttpClient:   http.DefaultClient,
	ApiKey:       "",
	ApiSecretKey: "",
})

func init() {
	logger.SetLevel(logger.DEBUG)
}

func TestByBitFutures_GetFutureDepth(t *testing.T) {
	t.Log(baDapi.GetFutureDepth(goex.ETH_USD, goex.QUARTER_CONTRACT, 10))
}

func TestByBitSwap_GetFutureTicker(t *testing.T) {
	ticker, err := baDapi.GetFutureTicker(goex.LTC_USD, goex.SWAP_CONTRACT)
	t.Log(err)
	t.Logf("%+v", ticker)
}

func TestByBit_GetExchangeInfo(t *testing.T) {
	baDapi.GetExchangeInfo()
}

func TestByBitFutures_GetFutureUserinfo(t *testing.T) {
	t.Log(baDapi.GetFutureUserinfo())
}

func TestByBitFutures_PlaceFutureOrder(t *testing.T) {
	//1044675677
	t.Log(baDapi.PlaceFutureOrder(goex.BTC_USD, goex.QUARTER_CONTRACT, "19990", "2", goex.OPEN_SELL, 0, 10))
}

func TestByBitFutures_LimitFuturesOrder(t *testing.T) {
	t.Log(baDapi.LimitFuturesOrder(goex.BTC_USD, goex.QUARTER_CONTRACT, "20001", "2", goex.OPEN_SELL))
}

func TestByBitFutures_MarketFuturesOrder(t *testing.T) {
	t.Log(baDapi.MarketFuturesOrder(goex.BTC_USD, goex.QUARTER_CONTRACT, "2", goex.OPEN_SELL))
}

func TestByBitFutures_GetFutureOrder(t *testing.T) {
	t.Log(baDapi.GetFutureOrder("1045208666", goex.BTC_USD, goex.QUARTER_CONTRACT))
}

func TestByBitFutures_FutureCancelOrder(t *testing.T) {
	t.Log(baDapi.FutureCancelOrder(goex.BTC_USD, goex.QUARTER_CONTRACT, "1045328328"))
}

func TestByBitFutures_GetFuturePosition(t *testing.T) {
	t.Log(baDapi.GetFuturePosition(goex.BTC_USD, goex.QUARTER_CONTRACT))
}

func TestByBitFutures_GetUnfinishFutureOrders(t *testing.T) {
	t.Log(baDapi.GetUnfinishFutureOrders(goex.BTC_USD, goex.QUARTER_CONTRACT))
}
