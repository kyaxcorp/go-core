package scalar

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"time"
)

func MarshalTimeDefault(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte("\"" + t.String() + "\""))
	})
}

func UnmarshalTimeDefault(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		// this is the format which is been exported...
		return time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", v)
	default:
		return time.Time{}, fmt.Errorf("%T is not a string", v)
	}
}
