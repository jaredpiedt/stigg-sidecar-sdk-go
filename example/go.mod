module github.com/jaredpiedt/stigg-sidecar-sdk-go/example

go 1.22.0

// Use version at HEAD, not the latest published.
replace github.com/jaredpiedt/stigg-sidecar-sdk-go => ../

require github.com/jaredpiedt/stigg-sidecar-sdk-go v0.0.0-00010101000000-000000000000

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
