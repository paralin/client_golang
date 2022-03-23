module github.com/prometheus/client_golang

go 1.16

replace (
	github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200726220404-fa7f51c52df0 // aperture-1.3.x
	google.golang.org/grpc => github.com/paralin/grpc-go v1.30.1-0.20210804030014-1587a7c16b66 // aperture
)

require (
	github.com/beorn7/perks v1.0.1
	github.com/cespare/xxhash/v2 v2.1.2
	github.com/golang/protobuf v1.5.2
	github.com/json-iterator/go v1.1.12
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.32.1
	github.com/prometheus/procfs v0.7.3
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9
	google.golang.org/protobuf v1.26.0
)
