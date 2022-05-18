package bybit

import (
	"net"
	"net/http"
	"net/url"
	"testing"
	"time"

	goex "github.com/nntaoli-project/goex"
)

var bs = NewByBitSwap(&goex.APIConfig{
	Endpoint: "https://testnet.binancefuture.com",
	HttpClient: &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse("socks5://127.0.0.1:1080")
			},
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
		},
		Timeout: 10 * time.Second,
	},
	ApiKey:       "",
	ApiSecretKey: "",
})

func TestByBitSwap_Ping(t *testing.T) {
	bs.Ping()
}

func TestByBitSwap_GetFutureDepth(t *testing.T) {
	t.Log(bs.GetFutureDepth(goex.BTC_USDT, "", 1))
}

func TestByBitSwap_GetFutureIndex(t *testing.T) {
	t.Log(bs.GetFutureIndex(goex.BTC_USDT))
}

func TestByBitSwap_GetKlineRecords(t *testing.T) {
	kline, err := bs.GetKlineRecords("", goex.BTC_USDT, goex.KLINE_PERIOD_4H, 1, goex.OptionalParameter{"test": 0})
	t.Log(err, kline[0].Kline)
}

func TestByBitSwap_GetTrades(t *testing.T) {
	t.Log(bs.GetTrades("", goex.BTC_USDT, 0))
}

func TestByBitSwap_GetFutureUserinfo(t *testing.T) {
	t.Log(bs.GetFutureUserinfo())
}

func TestByBitSwap_PlaceFutureOrder(t *testing.T) {
	t.Log(bs.PlaceFutureOrder(goex.BTC_USDT, "", "8322", "0.01", goex.OPEN_BUY, 0, 0))
}

func TestByBitSwap_PlaceFutureOrder2(t *testing.T) {
	t.Log(bs.PlaceFutureOrder(goex.BTC_USDT, "", "8322", "0.01", goex.OPEN_BUY, 1, 0))
}

func TestByBitSwap_GetFutureOrder(t *testing.T) {
	t.Log(bs.GetFutureOrder("1431689723", goex.BTC_USDT, ""))
}

func TestByBitSwap_FutureCancelOrder(t *testing.T) {
	t.Log(bs.FutureCancelOrder(goex.BTC_USDT, "", "1431554165"))
}

func TestByBitSwap_GetFuturePosition(t *testing.T) {
	t.Log(bs.GetFuturePosition(goex.BTC_USDT, ""))
}
