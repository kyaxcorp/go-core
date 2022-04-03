package client

import "github.com/rs/zerolog"

/*
We should enable logging throughout the entire client, that will help us do a better debug

1. Add Enable/Disable logging
2. Add different log levels
3. Add a name of the log instance... usually it should take from the config or somehow?!... or we can setup also manually a name
4.
*/

// LDebug -> 0
func (c *Client) LDebug() *zerolog.Event {
	return c.Logger.Debug()
}

// LInfo -> 1
func (c *Client) LInfo() *zerolog.Event {
	return c.Logger.Info()
}

// LWarn -> 2
func (c *Client) LWarn() *zerolog.Event {
	return c.Logger.Warn()
}

// LError -> 3
func (c *Client) LError() *zerolog.Event {
	return c.Logger.Error()
}

// LFatal -> 4
func (c *Client) LFatal() *zerolog.Event {
	return c.Logger.Fatal()
}

// LPanic -> 5
func (c *Client) LPanic() *zerolog.Event {
	return c.Logger.Panic()
}

//

//-------------------------------------\\

//

func (c *Client) LEvent(eventType string, eventName string, beforeMsg func(event *zerolog.Event)) {
	c.Logger.InfoEvent(eventType, eventName, beforeMsg)
}

func (c *Client) LEventCustom(eventType string, eventName string) *zerolog.Event {
	return c.Logger.InfoEventCustom(eventType, eventName)
}

func (c *Client) LEventF(eventType string, eventName string, functionName string) *zerolog.Event {
	return c.Logger.InfoEventF(eventType, eventName, functionName)
}

//

//-------------------------------------\\

//

// LWarnF -> when you need specifically to indicate in what function the logging is happening
func (c *Client) LWarnF(functionName string) *zerolog.Event {
	return c.Logger.WarnF(functionName)
}

// LInfoF -> when you need specifically to indicate in what function the logging is happening
func (c *Client) LInfoF(functionName string) *zerolog.Event {
	return c.Logger.InfoF(functionName)
}

// LDebugF -> when you need specifically to indicate in what function the logging is happening
func (c *Client) LDebugF(functionName string) *zerolog.Event {
	return c.Logger.DebugF(functionName)
}

// LErrorF -> when you need specifically to indicate in what function the logging is happening
func (c *Client) LErrorF(functionName string) *zerolog.Event {
	return c.Logger.ErrorF(functionName)
}
