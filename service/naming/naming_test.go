package naming

import (
	"testing"

	dpb "github.com/binchencoder/gateway-proto/data"
)

func TestToString(t *testing.T) {
	vexillaryServiceName, err := ServiceIdToName(dpb.ServiceId_VEXILLARY_SERVICE)
	if err != nil {
		t.Error("Expect no error")
	}
	if vexillaryServiceName != "vexillary-service" {
		t.Errorf("Expect 'vexillary-service' but got %s.", vexillaryServiceName)
	}

	vexillaryServiceId, err := ServiceNameToId(vexillaryServiceName)
	if err != nil {
		t.Error("Expect no error")
	}
	if vexillaryServiceId != dpb.ServiceId_VEXILLARY_SERVICE {
		t.Errorf("Expect 'ServiceId_VEXILLARY_SERVICE' but got %v.", vexillaryServiceId)
	}

	vexillaryServiceName, err = ServiceIdToName(dpb.ServiceId(13131313)) // Some non-existent service id.
	if err == nil {
		t.Error("Expect error")
	}

	vexillaryServiceId, err = ServiceNameToId("wrong-name")
	if err == nil {
		t.Error("Expect error")
	}
}
