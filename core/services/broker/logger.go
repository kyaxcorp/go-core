package broker

import "github.com/rs/zerolog"

/*
We should enable logging throughout the entire client, that will help us do a better debug

1. Add Enable/Disable logging
2. Add different log levels
3. Add a name of the log instance... usually it should take from the config or somehow?!... or we can setup also manually a name
4.
*/

// LDebug -> 0
func (b *Broker) LDebug() *zerolog.Event {
	return b.Logger.Debug()
}

// LInfo -> 1
func (b *Broker) LInfo() *zerolog.Event {
	return b.Logger.Info()
}

// LWarn -> 2
func (b *Broker) LWarn() *zerolog.Event {
	return b.Logger.Warn()
}

// LError -> 3
func (b *Broker) LError() *zerolog.Event {
	return b.Logger.Error()
}

// LFatal -> 4
func (b *Broker) LFatal() *zerolog.Event {
	return b.Logger.Fatal()
}

// LPanic -> 5
func (b *Broker) LPanic() *zerolog.Event {
	return b.Logger.Panic()
}

//

//-------------------------------------\\

//

func (b *Broker) LEvent(eventType string, eventName string, beforeMsg func(event *zerolog.Event)) {
	b.Logger.InfoEvent(eventType, eventName, beforeMsg)
}

func (b *Broker) LEventCustom(eventType string, eventName string) *zerolog.Event {
	return b.Logger.InfoEventCustom(eventType, eventName)
}

func (b *Broker) LEventF(eventType string, eventName string, functionName string) *zerolog.Event {
	return b.Logger.InfoEventF(eventType, eventName, functionName)
}

//

//-------------------------------------\\

//

// LWarnF -> when you need specifically to indicate in what function the logging is happening
func (b *Broker) LWarnF(functionName string) *zerolog.Event {
	return b.Logger.WarnF(functionName)
}

// LInfoF -> when you need specifically to indicate in what function the logging is happening
func (b *Broker) LInfoF(functionName string) *zerolog.Event {
	return b.Logger.InfoF(functionName)
}

// LDebugF -> when you need specifically to indicate in what function the logging is happening
func (b *Broker) LDebugF(functionName string) *zerolog.Event {
	return b.Logger.DebugF(functionName)
}

// LErrorF -> when you need specifically to indicate in what function the logging is happening
func (b *Broker) LErrorF(functionName string) *zerolog.Event {
	return b.Logger.ErrorF(functionName)
}
