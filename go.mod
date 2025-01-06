module github.com/razvanmarinn/tinyman-go-sdk

go 1.23

toolchain go1.23.4

require (
	github.com/algorand/go-algorand-sdk/v2 v2.7.0
	github.com/synycboom/tinyman-go-sdk v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.29.0
)

require (
	github.com/algorand/avm-abi v0.2.0 // indirect
	github.com/algorand/go-codec/codec v1.1.10 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
)

replace github.com/synycboom/tinyman-go-sdk => github.com/razvanmarinn/tinyman-go-sdk v0.0.0-20250106231111-2f83cd5b53ca
