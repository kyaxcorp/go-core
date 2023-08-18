package autoloader

var beforeLoadSetDefaults = make(map[string]func())

func BeforeLoadSetDefaults(name string, cb func()) {
	beforeLoadSetDefaults[name] = cb
}
