# Stigg Sidecar Go SDK

A Go client library for accessing the Stigg [Sidecar](https://docs.stigg.io/docs/sidecar).

## Installation

```bash
go get github.com/jaredpiedt/stigg-sidecar-sdk-go
```

## Usage

```go
package main

import (
	"context"

	"github.com/jaredpiedt/stigg-sidecar-sdk-go"
)

func main() {
	ctx := context.Background()

	client, err := stigg.NewClient("localhost:8443")
	if err != nil {
		log.Fatalf(err)
	}
    defer client.Close()

	resp, err := client.GetEntitlements(ctx, &stigg.GetEntitlementsRequest{
		CustomerId: *customerID,
	})
	if err != nil {
		log.Fatalf(err)
	}
}
```
