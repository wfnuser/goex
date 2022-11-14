package common

import "github.com/nntaoli-project/goex/v2"

func AdaptKlinePeriodToSymbol(period goex.KlinePeriod) string {
	switch period {
	case goex.Kline_1min:
		return "1m"
	case goex.Kline_5min:
		return "5m"
	case goex.Kline_15min:
		return "15m"
	case goex.Kline_30min:
		return "30m"
	case goex.Kline_60min, goex.Kline_1h:
		return "1H"
	case goex.Kline_4h:
		return "4H"
	case goex.Kline_6h:
		return "6H"
	case goex.Kline_1day:
		return "1D"
	case goex.Kline_1week:
		return "1W"
	default:
		return string(period)
	}
}

func adaptOrderSideToSym(s goex.OrderSide) string {
	switch s {
	case goex.Spot_Buy:
		return "buy"
	case goex.Spot_Sell:
		return "sell"
	}
	return ""
}

func adaptOrderTypeToSym(ty goex.OrderType) string {
	switch ty {
	case goex.OrderType_Limit:
		return "limit"
	case goex.OrderType_Market:
		return "market"
	}
	return string(ty.String())
}
