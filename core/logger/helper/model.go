package helper

import (
	"github.com/kyaxcorp/go-core/core/logger/channel"
	"github.com/kyaxcorp/go-core/core/logger/model"
	"github.com/rs/zerolog"
)

/*
This helper may be used when you need to have a Logger inside your existing model without calling
many times Logger.Logger...
*/

type Logger struct {
	// Logger -> you can set the logger by yourself, or you can indicate a channel name
	// and it will bind automatically with it
	Logger *model.Logger

	// ChannelName -> it's automatically binding with the config file if exists!
	ChannelName string
	// These are additional info/params
	ModuleName    string
	SubModuleName string
	FunctionName  string
	//
	VersionNr string

	// These are additional keys
	AdditionalInfo      map[string]string
	isAdditionalInfoSet bool
}

func New(l *Logger) (*Logger, error) {
	// If logger is nil, then create it
	if l == nil {
		l = &Logger{}
	}
	// Create the map
	l.AdditionalInfo = make(map[string]string)

	// Check if the channel has been set
	if l.ChannelName != "" {
		c, _err := channel.GetChannel(channel.Config{
			ChannelName:              l.ChannelName,
			ReturnDefaultIfNotExists: true,
		})
		if _err != nil {
			return nil, _err
		}
		// Set the Logger
		l.Logger = c
	} else {
		// Check if the Logger has been set!
		// If not then set the default one!

		if l.Logger == nil {
			c, _err := channel.GetDefaultChannel()
			if _err != nil {
				return nil, _err
			}
			// Set the Logger
			l.Logger = c
		} else {
			// has been set!
		}
	}

	return l, nil
}
func (l *Logger) SetModuleName(moduleName string) *Logger {
	l.ModuleName = moduleName
	return l
}

func (l *Logger) SetSubModuleName(subModuleName string) *Logger {
	l.SubModuleName = subModuleName
	return l
}

func (l *Logger) SetFunctionName(functionName string) *Logger {
	l.FunctionName = functionName
	return l
}

func (l *Logger) SetVersionNr(versionNr string) *Logger {
	l.VersionNr = versionNr
	return l
}

func (l *Logger) SetAddInfo(k, v string) *Logger {
	l.AdditionalInfo[k] = v
	l.isAdditionalInfoSet = true
	return l
}

func (l *Logger) AddAdditionalInfo(e *zerolog.Event) *zerolog.Event {
	if l.ModuleName != "" {
		e.Str("module_name", l.ModuleName)
	}
	if l.SubModuleName != "" {
		e.Str("sub_module_name", l.SubModuleName)
	}
	if l.FunctionName != "" {
		e.Str("function_name", l.FunctionName)
	}
	if l.VersionNr != "" {
		e.Str("version_nr", l.VersionNr)
	}
	if l.isAdditionalInfoSet {
		for k, v := range l.AdditionalInfo {
			e.Str(k, v)
		}
	}

	return e
}

// LDebug -> 0
func (l *Logger) LDebug() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Debug())
}

// LInfo -> 1
func (l *Logger) LInfo() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Info())
}

// LWarn -> 2
func (l *Logger) LWarn() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Warn())
}

// LError -> 3
func (l *Logger) LError() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Error())
}

// LFatal -> 4
func (l *Logger) LFatal() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Fatal())
}

// LPanic -> 5
func (l *Logger) LPanic() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Panic())
}

//

//

//-------------------------------------\\

// LD -> 0
func (l *Logger) LD() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Debug())
}

// LI -> 1
func (l *Logger) LI() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Info())
}

// LW -> 2
func (l *Logger) LW() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Warn())
}

// LE -> 3
func (l *Logger) LE() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Error())
}

// LF -> 4
func (l *Logger) LF() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Fatal())
}

// LP -> 5
func (l *Logger) LP() *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.Panic())
}

//

//-------------------------------------\\

//

func (l *Logger) LEvent(eventType string, eventName string, beforeMsg func(event *zerolog.Event)) {
	l.Logger.InfoEvent(eventType, eventName, beforeMsg)
}

func (l *Logger) LEventCustom(eventType string, eventName string) *zerolog.Event {
	return l.Logger.InfoEventCustom(eventType, eventName)
}

func (l *Logger) LEventF(eventType string, eventName string, functionName string) *zerolog.Event {
	return l.Logger.InfoEventF(eventType, eventName, functionName)
}

//

//-------------------------------------\\

//

// LWarnF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LWarnF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.WarnF(functionName))
}

// LInfoF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LInfoF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.InfoF(functionName))
}

// LDebugF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LDebugF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.DebugF(functionName))
}

// LErrorF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LErrorF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.ErrorF(functionName))
}

// LFatalF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LFatalF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.FatalF(functionName))
}

// LPanicF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LPanicF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.PanicF(functionName))
}

//

//

// LWF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LWF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.WarnF(functionName))
}

// LIF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LIF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.InfoF(functionName))
}

// LDF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LDF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.DebugF(functionName))
}

// LEF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LEF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.ErrorF(functionName))
}

// LFF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LFF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.FatalF(functionName))
}

// LPF -> when you need specifically to indicate in what function the logging is happening
func (l *Logger) LPF(functionName string) *zerolog.Event {
	return l.AddAdditionalInfo(l.Logger.PanicF(functionName))
}
