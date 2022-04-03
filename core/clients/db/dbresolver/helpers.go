package dbresolver

import (
	"github.com/gookit/color"
	"github.com/kyaxcorp/go-core/core/helpers/conv"
	"github.com/rs/zerolog"
	"time"
)

func printConsumedTime(info func() *zerolog.Event, funcStartTime int) {
	funcEndTime := time.Now().Nanosecond()
	totalConsumedTime := funcEndTime - funcStartTime
	info().
		Int("consumed_time", totalConsumedTime).
		Msg(color.LightBlue.Render("consumed time -> Nano:" + conv.IntToStr(totalConsumedTime)))
}
