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

	DefaultParameterStart = 0
	DefaultParameterLimit = 1000
)

const (
	DefaultKubeHost   = ""
	DefaultKubeConfig = ""
	DefaultListenPort = 80

	DefaultDatabaseString = "/caicloud/simple-object-storage/db"
	DefaultStorageString  = "/caicloud/simple-object-storage/mnt/glusterfs-single"
)
