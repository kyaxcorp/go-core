package channels

import (
	"log"
	"time"
)

func main() {
	/*var jordan map[string]map[uint64]string
	jordan = make(map[string]map[uint64]string)
	if _, ok := jordan["vasea"]; !ok {
		jordan["vasea"] = make(map[uint64]string)
	}
	jordan["vasea"][232323] = "hello"
	log.Println(jordan)
	log.Println(jordan["vasea"][232323])*/

	mainChannel := make(chan uint)

	go func() {
		for {
			select {
			case channelData := <-mainChannel:
				log.Println("Creating", channelData)
				go func() {
					time.Sleep(time.Second * 3)
					log.Println(channelData)
				}()
			}
		}
	}()

	i := 0
	for {
		i++
		// Send data!
		mainChannel <- uint(i)
		time.Sleep(time.Second)
	}
}
