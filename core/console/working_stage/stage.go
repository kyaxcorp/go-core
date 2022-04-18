package working_stage

var isDev bool

func IsDev() bool {
	return isDev
}

func IsProd() bool {
	return !isDev
}

func SetStage(isDev bool) {
	isDev = isDev
}
