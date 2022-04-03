package multi_writer

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/conv"
	"github.com/gookit/color"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"io"
)

// This is a writer which has a method Write and returns the interface of io.Writer
type multiWriter struct {
	//writers []io.Writer
	writers []CustomWriter
}

type CustomWriter struct {
	Writer io.Writer
	// Should colors be filtered for this writer
	FilterColors bool
	// Here we define what type of writer is this, like stdout, file etc...
	Type string
}

// MultiWriter -> create the instance which will handle the writing
func MultiWriter(multiWriters []CustomWriter) io.Writer {
	return &multiWriter{
		writers: multiWriters,
	}
}

// Write -> This is the main function which will handle the writing
func (t *multiWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		// Remove stdout colors
		if w.FilterColors {
			// Convert from bytes to string
			msg := conv.BytesToStr(p)
			// Read the json value
			_msg := gjson.Get(msg, "message").String()
			// Remove the color codes
			_msg = color.ClearCode(_msg)
			//log.Println("found value", _msg)
			// Set back to json the message value
			newMsg, _err := sjson.Set(msg, "message", _msg)
			if _err != nil {
				// Take the original...
				newMsg = msg
			}
			// convert back to bytes
			p = conv.StrToBytes(newMsg)
		}
		n, err = w.Writer.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}
