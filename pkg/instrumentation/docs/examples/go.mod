module github.com/Azure/ASC-go-libs/examples

go 1.16

replace github.com/Azure/ASC-go-libs/pkg/config => ./../../../config
replace github.com/Azure/ASC-go-libs/pkg/instrumentation => ./../..
replace github.com/Azure/ASC-go-libs/pkg/common => ./../../../common

require (
 github.com/Azure/ASC-go-libs/pkg/config v1.0.0
 github.com/Azure/ASC-go-libs/pkg/instrumentation v1.0.0
)
