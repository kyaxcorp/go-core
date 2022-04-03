package arrays

import "log"

func main() {
	var jordan map[string]map[uint64]string
	jordan = make(map[string]map[uint64]string)
	if _, ok := jordan["vasea"]; !ok {
		jordan["vasea"] = make(map[uint64]string)
	}
	jordan["vasea"][232323] = "hello"
	log.Println(jordan)
	log.Println(jordan["vasea"][232323])

}
