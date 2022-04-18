package working_stage

var IsDev bool

func GetStage() bool {
	return IsDev
}

func SetStage(isDev bool) {
	IsDev = isDev
}
