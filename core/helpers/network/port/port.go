package port

import (
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"net"
	"strings"
	"sync"
)

// add a lock when searching for a free address!
// because multiple goroutines can call this function, and a conflict can appear between them!
// For example if we are launching websocket and http at the same time, they can conflict between them!
// And when a free address has been found, we will add to busy one!
var searchFreeAddressLock sync.Mutex
var occupiedAddresses = make(map[string]bool)

// IsTCPBusy -> checks if port is busy!
func IsTCPBusy(address string) (bool, error) {
	return IsBusy("tcp", address)
}

func SearchAndLockFreeTCPAddress(address string) (string, error) {
	return SearchAndLockFreeAddress("tcp", address)
}

func SearchAndLockFreeUDPAddress(address string) (string, error) {
	return SearchAndLockFreeAddress("udp", address)
}

func SearchAndLockFreeAddress(protocol string, address string) (string, error) {
	searchFreeAddressLock.Lock()
	defer searchFreeAddressLock.Unlock()

	// Check if it has + as auto search
	//if !strings.Contains(address, "+") {
	//	return address, nil
	//}

	// Get the ip address
	ipAddress := ExtractIPAddress(address)
	if ipAddress == "" {
		return address, define.Err(0, "ip address is empty")
	}

	// The recheck will run until it will find a free address to bind to!
recheck:
	// Check if it's busy...
	isBusy := false
	if busy, _ := IsBusy(protocol, address); busy {
		isBusy = true
	}

	// Check also here:
	if !isBusy {
		if _, ok := occupiedAddresses[address]; ok {
			// it's also busy...
			isBusy = true
		}
	}

	if isBusy {
		// if it's busy get the address port... and do +1
		p := GetFreeAddressPort(address)
		if p == nil {
			// then there is no port defined...
			// so let's return the same address
			return address, define.Err(0, "port is not defined")
		}
		newPort := *p + 1
		if newPort >= 65535 {
			// Stop here!
			return "", define.Err(0, "no ports available, the max one is reached")
		}

		address = ipAddress + ":" + conv.IntToStr(newPort)
		// retry...
		goto recheck
	}

	// return same address if it's not used
	return address, nil
}

func ExtractIPAddress(address string) string {
	address = FilterAddress(address)
	if strings.Contains(address, ":") {
		s := strings.Split(address, ":")
		if len(s) >= 1 {
			return s[0]
		}
		// no address has been found
		return ""
	}
	// return same address if no : indicated
	return address
}

func GetFreeAddressPort(address string) *int {
	address = FilterAddress(address)
	if strings.Contains(address, ":") {
		s := strings.Split(address, ":")
		if len(s) <= 1 {
			return nil
		}
		p := conv.StrToInt(s[1])
		return &p
	}
	return nil
}

// IsUDPBusy -> checks if port is busy!
func IsUDPBusy(address string) (bool, error) {
	return IsBusy("udp", address)
}

func FilterAddress(address string) string {
	if strings.Contains(address, "+") {
		address = strings.ReplaceAll(address, "+", "")
	}
	return address
}

func IsBusy(protocol string, address string) (bool, error) {
	// Is Busy
	address = FilterAddress(address)
	ln, _err := net.Listen(protocol, address)
	if _err != nil {
		return true, _err
	}
	ln.Close()
	return false, nil
}
