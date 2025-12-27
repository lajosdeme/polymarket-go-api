package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lajosdeme/polymarket-go-api/api"
	"github.com/lajosdeme/polymarket-go-api/client"
	"github.com/lajosdeme/polymarket-go-api/types"
)

func testGamma() {
	gammaClient := client.NewGammaClient("")
	gammaAPI := api.NewGammaAPI(gammaClient)

	ctx := context.Background()
	fmt.Println("Testing Gamma API...")

	markets, err := gammaAPI.GetMarkets(ctx, &types.MarketFilters{
		Limit: gammaIntPtr(5),
	})
	if err != nil {
		log.Printf("Failed to get markets: %v", err)
		return
	}

	fmt.Printf("Got %d markets\n", len(markets))

	for _, m := range markets {
		fmt.Println(stringPtrValue(m.Question))
		fmt.Println(stringPtrValue(m.ClobTokenIds))
	}
}

func gammaIntPtr(i int) *int {
	return &i
}

func stringPtrValue(s *string) string {
	if s == nil {
		return "N/A"
	}
	return *s
}
