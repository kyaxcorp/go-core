package time

import (
	"time"
)

var savedLocalTimezone *time.Location

func OverrideLocalTimezone(timezone string) (bool, error) {

	location, _err := time.LoadLocation(timezone)
	// Save before overriding

	if location == nil || _err != nil {
		return false, _err
	}

	saveLocalTimeZone()

	// log.Println(location.String())
	time.Local = location
	return true, nil
}

func saveLocalTimeZone() {
	if savedLocalTimezone == nil {
		cloned := *time.Local
		savedLocalTimezone = &cloned
		//savedLocalTimezone = time.Local
	}
}

func OverrideLocalTimezoneByLocation(location *time.Location) (bool, error) {
	saveLocalTimeZone()
	// log.Println(location.String())
	time.Local = location
	return true, nil
}

func GetOriginalLocalTimezone() string {
	if savedLocalTimezone == nil {
		// it means it wasn't overrided...
		return time.Local.String()
	} else {
		return savedLocalTimezone.String()
	}
}

func RestoreOriginalLocalTimezone() {
	if savedLocalTimezone == nil {
		// it means it wasn't overrided...
	} else {
		OverrideLocalTimezoneByLocation(savedLocalTimezone)
	}
}

func GetOriginalLocalTimezoneLocation() *time.Location {
	if savedLocalTimezone == nil {
		// it means it wasn't overrided...
		return time.Local
	} else {
		return savedLocalTimezone
	}
}
