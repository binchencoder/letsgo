package naming

import (
	"fmt"
	"strings"

	dpb "binchencoder.com/gateway-proto/data"
)

// ServiceIdToName converts service id to service name in
// standard string form.
func ServiceIdToName(serviceId dpb.ServiceId) (string, error) {
	serviceName, ok := dpb.ServiceId_name[int32(serviceId)]
	if ok {
		return strings.ToLower(strings.Replace(serviceName, "_", "-", -1)), nil
	}
	return fmt.Sprintf("!%d", serviceId), fmt.Errorf("Invalid service id %d", serviceId)
}

// ServiceNameToId converts service name in standard string form
// to service id.
func ServiceNameToId(serviceName string) (dpb.ServiceId, error) {
	serviceId, ok := dpb.ServiceId_value[strings.ToUpper(strings.Replace(serviceName, "-", "_", -1))]
	if ok {
		return dpb.ServiceId(serviceId), nil
	}
	return dpb.ServiceId_SERVICE_NONE, fmt.Errorf("Invalid service name %s", serviceName)
}

// ServiceNameToFolderName converts service name to folder name used
// in production repo.
func ServiceNameToFolderName(serviceName string) string {
	return strings.ToLower(strings.Replace(serviceName, "_", "-", -1))
}

// FolderNameToServiceName converts folder name to service name.
func FolderNameToServiceName(folderName string) string {
	return strings.ToUpper(strings.Replace(folderName, "-", "_", -1))
}
