package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/jaredpiedt/stigg-sidecar-sdk-go"
)

var (
	addr       = flag.String("addr", "localhost:8443", "The address to connect to")
	customerID = flag.String("customerId", "", "The customer id to retrieve entitlements for")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	client, err := stigg.NewClient(*addr)
	if err != nil {
		log.Fatalf("could not create client: %v", err)
	}

	resp, err := client.GetEntitlements(ctx, &stigg.GetEntitlementsRequest{
		CustomerId: *customerID,
	})
	if err != nil {
		log.Fatalf("could not get entitlements: %v", err)
	}

	for _, e := range resp.Entitlements {
		if e.GetBoolean() != nil {
			fmt.Printf("BOOLEAN: %s\n", e.GetBoolean().GetFeature().GetId())
		}
		if e.GetMetered() != nil {
			fmt.Printf("METERED: %s\n", e.GetMetered().GetFeature().GetId())
		}
		if e.GetNumeric() != nil {
			fmt.Printf("NUMERIC: %s\n", e.GetNumeric().GetFeature().GetId())
		}
	}
}
