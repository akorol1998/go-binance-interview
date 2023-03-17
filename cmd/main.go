package main

import (
	"context"
	"fmt"
	"go-binance-interview/internal/handler"
	"sync"

	"github.com/aiviaio/go-binance/v2"
)

func main() {
	var wg sync.WaitGroup
	var (
		apiKey    = "apiKey"
		secretKey = "secretKey"
	)

	// Init the value here
	n := 5
	ch := make(chan *binance.SymbolPrice, n)
	client := handler.InitClient(apiKey, secretKey)
	if client == nil {
		panic("Faield to init the client")
	}
	ctx := context.Background()

	symb, err := client.ExchangeInfo(ctx)
	if err != nil {
		panic("Faield to fetch ExchangeInfo")
	}
	resSl := symb[:n]

	for _, symbol := range resSl {
		wg.Add(1)
		go func(s string, wg *sync.WaitGroup) {
			defer wg.Done()
			ctx := context.Background()
			res, err := client.SymbolPrice(ctx, s)
			if err != nil {
				panic("Faield to fetch Price for a symbol")
			}
			ch <- res[0]
		}(symbol.Symbol, &wg)
	}
	wg.Wait()
	close(ch)

loop:
	for {
		select {
		case val := <-ch:
			fmt.Println(val.Symbol, val.Price)
			if len(ch) == 0 {
				break loop
			}
		}
	}
}
