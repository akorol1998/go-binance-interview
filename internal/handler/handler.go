package handler

import (
	"context"

	"github.com/aiviaio/go-binance/v2"
)

type BClient struct {
	Client *binance.Client
}

func InitClient(apiKey, secretKey string) *BClient {
	return &BClient{binance.NewClient(apiKey, secretKey)}
}

// ExchangeInfo returns exchange info with no filters
// ctx: Context
func (bc *BClient) ExchangeInfo(ctx context.Context) ([]binance.Symbol, error) {
	res, err := bc.Client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, err
	}
	return res.Symbols, nil
}

// SymbolPrice returns most recent price for the symbol
// ctx: context
// symbol: symbol we are interested
func (bc *BClient) SymbolPrice(ctx context.Context, symbol string) ([]*binance.SymbolPrice, error) {
	res, err := bc.Client.NewListPricesService().Symbol(symbol).Do(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil

}
