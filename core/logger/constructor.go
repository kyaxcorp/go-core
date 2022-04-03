package logger

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/conv"
	"github.com/KyaXTeam/go-core/v2/core/helpers/filesystem"
	"github.com/KyaXTeam/go-core/v2/core/helpers/io/writer"
	"github.com/KyaXTeam/go-core/v2/core/logger/application/vars"
	"github.com/KyaXTeam/go-core/v2/core/logger/config"
	"github.com/KyaXTeam/go-core/v2/core/logger/model"
	"github.com/KyaXTeam/go-core/v2/core/logger/multi_writer"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

// GetAppLogger -> it returns the instance which is the main logger of the app, it centralizes all the data together
func GetAppLogger() *model.Logger {
	return vars.ApplicationLogger
}

//func New(ctx context.Context, config config.Config) *Logger {

// Here we will store writers which handle writing to file, if we don't want to create multiple handlers and have conflict,
// between them, it's better to have here unique handlers for each file
var writersByFileFullPath = make(map[string]io.Writer)

// This Lock is for doing operations on writersByFileFullPath
var writersByFileFullPathLock sync.Mutex

// We create and save 1 instance of console writer! Why? Because if having multiple, dangerous things like
// overlapped messages can occur!
var consoleWriter *zerolog.ConsoleWriter

// Here we store the main application writer instance
var applicationWriter *io.Writer

var conversionLevels = map[int]zerolog.Level{
	1: zerolog.TraceLevel,
	2: zerolog.DebugLevel,
	3: zerolog.InfoLevel,
	4: zerolog.WarnLevel,
	5: zerolog.ErrorLevel,
	6: zerolog.FatalLevel,
	7: zerolog.PanicLevel,
}

// New -> creates a new logger client
func New(config config.Config) *model.Logger {
	// The context
	//if ctx == nil {
	//	ctx = _context.GetRootContext()
	//}
	// The Writers
	var writers []multi_writer.CustomWriter

	// We have set this format because it doesn't show us milliseconds when setting this format 15:04:05.000
	zerolog.TimeFieldFormat = time.RFC3339Nano
	// Check if console writer has being created
	if consoleWriter == nil {
		// Create the console writer
		colorStdOut := colorable.NewColorableStdout()

		// Set the time format from the config!
		timeFormat := "002 15:04:05.000"
		if config.ConsoleTimeFormat != "" {
			timeFormat = config.ConsoleTimeFormat
		}
		consoleWriter = &zerolog.ConsoleWriter{
			Out:        colorStdOut,
			TimeFormat: timeFormat,
		}
	}

	// Is console logging is enabled
	if conv.ParseBool(config.IsEnabled) && conv.ParseBool(config.ConsoleIsEnabled) {
		//writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
		writers = append(writers, multi_writer.CustomWriter{
			Writer: *consoleWriter,
		})
	}

	// Is File logging is enabled
	var mainWriter io.Writer
	if conv.ParseBool(config.IsEnabled) && conv.ParseBool(config.FileIsEnabled) {
		mainWriter = getFileHandler(config)
		writers = append(writers, multi_writer.CustomWriter{
			Writer: mainWriter,
			// We don't need colors in file
			FilterColors: true,
		})
	}

	// If it's not the master app
	if conv.ParseBool(config.IsEnabled) && !conv.ParseBool(config.IsApplication) {
		// Add additional writer to master, take the ApplicationLogger Config
		// We should get an existent instance of Application writer
		// Don't know what will happen if multiple processes of the same application will work, but don't see
		// any problems here!

		if applicationWriter == nil {
			appWriter := getFileHandler(GetAppLogger().Config)
			applicationWriter = &appWriter
		}

		writers = append(writers, multi_writer.CustomWriter{
			Writer: *applicationWriter,
			// We don't need colors in file
			FilterColors: true,
		})
	}

	// Check if there is the main writer and if write to parent is enabled
	if conv.ParseBool(config.IsEnabled) &&
		config.ParentWriter != nil &&
		conv.ParseBool(config.WriteToParent) {
		// Add to writers parent's writer
		writers = append(writers, multi_writer.CustomWriter{
			Writer: config.ParentWriter,
			// We don't need colors in file
			FilterColors: true,
		})
	}

	//  Create the multi writers
	mw := multi_writer.MultiWriter(writers)

	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// Create the logger itself
	//logger := zerolog.New(mw).With().Timestamp().Logger().WithContext(ctx)

	logger := zerolog.New(mw).
		With().
		Str("module", config.ModuleName).
		Timestamp().
		Logger().
		Level(ConvertConfigLogLevel(config.Level))

	return &model.Logger{
		Config: config,
		Logger: &logger,
		// Set Main File Writer as reference!
		MainWriter: mainWriter,
		//parentCtx: ctx,
	}
}

func ConvertConfigLogLevel(level int) zerolog.Level {

	// panic (zerolog.PanicLevel, 5)
	// fatal (zerolog.FatalLevel, 4)
	// error (zerolog.ErrorLevel, 3)
	// warn (zerolog.WarnLevel, 2)
	// info (zerolog.InfoLevel, 1)
	// debug (zerolog.DebugLevel, 0)
	// trace (zerolog.TraceLevel, -1)

	if val, ok := conversionLevels[level]; ok {
		return val
	} else {
		// Return the default value if it's an indicated an incorrect one
		return zerolog.WarnLevel
	}
}

func getFileHandler(config config.Config) io.Writer {
	fileName := config.Name + ".log"
	if config.FileName != "" {
		fileName = config.FileName
	}
	fullFilePath := filesystem.FilterPath(path.Join(config.DirLogPath, fileName))
	writersByFileFullPathLock.Lock()
	defer writersByFileFullPathLock.Unlock()
	if _writer, ok := writersByFileFullPath[fullFilePath]; ok {
		return _writer
	} else {
		_writer := newRollingFile(config)
		// Save the writer
		writersByFileFullPath[fullFilePath] = _writer
		// Return in!
		return _writer
	}
}

func newRollingFile(config config.Config) io.Writer {
	// TODO: maybe this creation is useless because the logger auto creates...
	if _err := os.MkdirAll(config.DirLogPath, 0744); _err != nil {
		log.Error().Err(_err).Str("path", config.DirLogPath).Msg("can't create log directory")
		return nil
	}

	// Override the file name if there is one...
	fileName := config.Name + ".log"
	if config.FileName != "" {
		fileName = config.FileName
	}

	return &writer.Logger{
		Filename:   filesystem.FilterPath(path.Join(config.DirLogPath, fileName)),
		MaxBackups: config.FileRotateMaxFiles, // files
		MaxSize:    config.FileLogMaxSize,     // megabytes
		MaxAge:     config.FileMaxAge,         // days
	}
}
