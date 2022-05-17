package openapi

import (
	"time"
)

// GetHwComponent -
func US_GetHwComponent() HwComponentInfo {
	hwComponent := HwComponentInfo{
		DeviceType: "transducer",
		Id:         "GHJGJHGJHGJHJGJHG",
	}
	return hwComponent
}

// Register -
func US_Register() HwComponentRegistrationInfo {
	regInfo := HwComponentRegistrationInfo{
		Info: HwComponentInfo{
			DeviceType: "some-device",
			Id:         "newid-DSAFADSEERW",
		},
		LastUpdate: time.Now(),
	}
	return regInfo
}
