package gor

import "github.com/rs/zerolog"

/*
We should enable logging throughout the entire client, that will help us do a better debug

1. Add Enable/Disable logging
2. Add different log levels
3. Add a name of the log instance... usually it should take from the config or somehow?!... or we can setup also manually a name
4.
*/

/*
func filter(logger *zerolog.Event) *zerolog.Event {
	// TODO: change key and name!
	return logger.Str("module", "goroutine")
}

// LDebug -> 0
func (g *GInstance) LDebug() *zerolog.Event {
	return filter(appLog.Debug())
}

// LInfo -> 1
func (g *GInstance) LInfo() *zerolog.Event {
	return filter(appLog.Info())
}

// LWarn -> 2
func (g *GInstance) LWarn() *zerolog.Event {
	return filter(appLog.Warn())
}

// LError -> 3
func (g *GInstance) LError() *zerolog.Event {
	return filter(appLog.Error())
}

// LFatal -> 4
func (g *GInstance) LFatal() *zerolog.Event {
	return filter(appLog.Fatal())
}

// LPanic -> 5
func (g *GInstance) LPanic() *zerolog.Event {
	return filter(appLog.Panic())
}

//

//-------------------------------------\\

//

// LWarnF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LWarnF(functionName string) *zerolog.Event {
	return filter(appLog.WarnF(functionName))
}

// LInfoF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LInfoF(functionName string) *zerolog.Event {
	return filter(appLog.InfoF(functionName))
}

// LDebugF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LDebugF(functionName string) *zerolog.Event {
	return filter(appLog.DebugF(functionName))
}

// LErrorF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LErrorF(functionName string) *zerolog.Event {
	return filter(appLog.ErrorF(functionName))
}*/

/*
We should enable logging throughout the entire client, that will help us do a better debug

1. Add Enable/Disable logging
2. Add different log levels
3. Add a name of the log instance... usually it should take from the config or somehow?!... or we can setup also manually a name
4.
*/

// LDebug -> 0
func (g *GInstance) LDebug() *zerolog.Event {
	return g.Logger.Debug()
}

// LInfo -> 1
func (g *GInstance) LInfo() *zerolog.Event {
	return g.Logger.Info()
}

// LWarn -> 2
func (g *GInstance) LWarn() *zerolog.Event {
	return g.Logger.Warn()
}

// LError -> 3
func (g *GInstance) LError() *zerolog.Event {
	return g.Logger.Error()
}

// LFatal -> 4
func (g *GInstance) LFatal() *zerolog.Event {
	return g.Logger.Fatal()
}

// LPanic -> 5
func (g *GInstance) LPanic() *zerolog.Event {
	return g.Logger.Panic()
}

//

//-------------------------------------\\

//

func (g *GInstance) LEvent(eventType string, eventName string, beforeMsg func(event *zerolog.Event)) {
	g.Logger.InfoEvent(eventType, eventName, beforeMsg)
}

func (g *GInstance) LEventCustom(eventType string, eventName string) *zerolog.Event {
	return g.Logger.InfoEventCustom(eventType, eventName)
}

func (g *GInstance) LEventF(eventType string, eventName string, functionName string) *zerolog.Event {
	return g.Logger.InfoEventF(eventType, eventName, functionName)
}

//

//-------------------------------------\\

//

// LWarnF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LWarnF(functionName string) *zerolog.Event {
	return g.Logger.WarnF(functionName)
}

// LInfoF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LInfoF(functionName string) *zerolog.Event {
	return g.Logger.InfoF(functionName)
}

// LDebugF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LDebugF(functionName string) *zerolog.Event {
	return g.Logger.DebugF(functionName)
}

// LErrorF -> when you need specifically to indicate in what function the logging is happening
func (g *GInstance) LErrorF(functionName string) *zerolog.Event {
	return g.Logger.ErrorF(functionName)
}
