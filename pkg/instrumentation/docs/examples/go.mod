module github.com/Azure/Tivan-Libs/examples

go 1.16

replace github.com/Azure/Tivan-Libs/pkg/config => ./../../../config
replace github.com/Azure/Tivan-Libs/pkg/instrumentation => ./../..
replace github.com/Azure/Tivan-Libs/pkg/common => ./../../../common

require (
 github.com/Azure/Tivan-Libs/pkg/config v1.0.0
 github.com/Azure/Tivan-Libs/pkg/instrumentation v1.0.0
)
