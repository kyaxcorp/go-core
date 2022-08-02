package scalar

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"time"
)

func MarshalTimeRFC3339Nano(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte("\"" + t.Format(time.RFC3339Nano) + "\""))
	})
}

func UnmarshalTimeRFC3339Nano(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339Nano, v)
	default:
		return time.Time{}, fmt.Errorf("%T is not a string", v)
	}
}
