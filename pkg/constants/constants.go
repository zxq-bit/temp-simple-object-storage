package constants

import (
	"fmt"
)

const (
	APIVersion = "v1alpha1"
)

var (
	RootPath = fmt.Sprintf("/api/%s", APIVersion)
)

const (
	ParameterStart = "start"
	ParameterLimit = "limit"

	ParameterRequestBody = "req"
	ParameterXUser       = "X-User"
	ParameterXTenant     = "X-Tenant"
)

const (
	DefaultKubeHost   = ""
	DefaultKubeConfig = ""
	DefaultListenPort = 80

	DefaultDbString = ""

	DefaultLocalStorageBucketNum = 100

	FileTimestampFormat = "20060102150405" // "20060102150405.000" // for js do not accept '.' in value
)
