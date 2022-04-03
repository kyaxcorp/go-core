package main

import (
	timezone "github.com/kyaxcorp/go-core/core/helpers/time"
	"log"
	"time"
)

func main() {
	//os.Setenv("TZ", "Africa/Cairo")
	/*location, _ := time.LoadLocation("UTC")
	log.Println(location.String())
	time.Local = location*/

	//log.Println(time.Now().UTC().Format("2006-01-02 15:04:05"))
	log.Println(timezone.GetOriginalLocalTimezone())

	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	timezone.OverrideLocalTimezone("UTC")
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	timezone.OverrideLocalTimezone("Europe/Chisinau")
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	timezone.OverrideLocalTimezone("Europe/Berlin")
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
	log.Println(timezone.GetOriginalLocalTimezone())
	timezone.RestoreOriginalLocalTimezone()
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
}
