module github.com/prometheus/client_golang

go 1.15

replace github.com/prometheus/procfs => github.com/paralin/prometheus_procfs v0.7.2-0.20210804014342-14dd3b9e4161 // gopherjs-compat

require (
	github.com/beorn7/perks v1.0.1
	github.com/cespare/xxhash/v2 v2.1.1
	github.com/golang/protobuf v1.4.3
	github.com/json-iterator/go v1.1.11
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.29.0
	github.com/prometheus/procfs v0.6.0
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40
)
