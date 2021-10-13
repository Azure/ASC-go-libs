module github.com/Azure/ASC-go-libs/pkg/instrumentation

go 1.16

replace github.com/Azure/ASC-go-libs/pkg/common => ../common

require (
	github.com/Azure/ASC-go-libs/pkg/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
