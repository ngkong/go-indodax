package indodax

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// These are url to open data for public. It doesn't need an API key to call these methods. You can call
	// simple GET request or open it directly from the browser.
	UrlPublic = "https://indodax.com/api"

	// To use Private API first you need to obtain your API credentials by logging into your indodax.com
	// account and open https://indodax.com/trade_api. These credentials contain "API Key" and "Secret
	// Key". Please keep these credentials safe.
	UrlPrivate = "https://indodax.com/tapi"

	// path connection to public api
	pathTicker    = "/%s/ticker"
	pathDepth     = "/%s/depth"
	pathSummaries = "/summaries"
	pathTrades    = "/%s/trades"
)

const (
	fieldNameAmount         = "amount"
	fieldNameWalletAddress  = "address"
	fieldNameBalance        = "balance"
	fieldNameBalanceHold    = "balance_hold"
	fieldNameHigh           = "high"
	fieldNameLow            = "low"
	fieldNameLast           = "last"
	fieldNameBuy            = "buy"
	fieldNameSell           = "sell"
	fieldNameTID            = "tid"
	fieldNameType           = "type"
	fieldNamePrice          = "price"
	fieldNameDate           = "date"
	fieldNameName           = "name"
	fieldNameTickers        = "tickers"
	fieldNamePrices24h      = "prices_24h"
	fieldNamePrices7d       = "prices_7d"
	fieldNameUserId         = "user_id"
	fieldNameProfilePicture = "profile_picture"
	fieldNameUserName       = "name"
	fieldNameUserServerTime = "server_time"
	fieldNameEmail          = "email"
)

var (
	debug = "PROD"
)

var (
	// ErrUnauthenticated define an error when user did not provide token
	// and secret keys when accessing private APIs.
	ErrUnauthenticated = fmt.Errorf("unauthenticated connection")

	// ErrInvalidPairName define an error if user call API with empty,
	// invalid or unknown pair's name.
	ErrInvalidPairName = fmt.Errorf("invalid or empty pair name")
)

const (
	apiTrade                  = "trade"
	apiTradeCancelOrder       = "cancelOrder"
	apiViewGetInfo            = "getInfo"
	apiViewGetOrder           = "getOrder"
	apiViewOpenOrders         = "openOrders"
	apiViewOrderHistory       = "orderHistory"
	apiViewTradeHistory       = "tradeHistory"
	apiViewTransactionHistory = "transHistory"
	apiWithdraw               = "withdrawCoin"
)

func SetDebug(active bool) string {

	if active {
		debug = "DEV"
	}
	return debug
}

func printDebug(info interface{}) {
	if debug == "DEV" {
		fmt.Printf("DEBUG >>> %v", info)
	}
}

func jsonToMapStringFloat64(in map[string]interface{}) (
	out map[string]float64, err error,
) {
	out = make(map[string]float64, len(in))

	for k, v := range in {
		val64, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
		if err != nil {
			return nil, err
		}
		k = strings.ToLower(k)
		out[k] = val64
	}
	return out, nil
}

func jsonToMapStringString(in map[string]interface{}) (
	out map[string]string, err error,
) {
	out = make(map[string]string, len(in))

	for k, v := range in {
		k = strings.ToLower(k)
		out[k] = fmt.Sprintf("%v", v)
	}
	return out, nil
}

//
// timestamp return current time in milliseconds as integer.
//
func timestamp() int32 {
	return int32(time.Now().Unix())
}

//
// timestampAsString return current time in milliseconds as string.
//
func timestampAsString() string {
	return fmt.Sprintf("%d", timestamp())
}